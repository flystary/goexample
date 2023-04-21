package main

import (
	"github.com/streadway/amqp"
	"log"
)

// 检查每个amqp调用的返回值
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 连接到RabbitMQ服务器
	conn, err := amqp.Dial("amqp://admin:123456@192.168.10.10:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// API都是用过该通道操作的
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明消息要发送到的队列
	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	// 将消息发布到声明的队列
	body := "Hello World!"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf("[x] Sent %s", body)

}
