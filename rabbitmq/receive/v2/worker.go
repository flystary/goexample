package main

import (
	"bytes"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://admin:123456@192.168.10.10:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		false, // // 注意这里传false,关闭自动消息确认
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	// 开启循环不断地消费消息
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Receied a message: %s", d.Body)
			count := bytes.Count(d.Body, []byte("."))
			t := time.Duration(count)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false) //手动传递消息
		}
	}()

	log.Printf("[*] Wating for messages. To exit press CTRL+C")
	<-forever
}
