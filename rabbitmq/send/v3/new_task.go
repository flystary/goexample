package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
)

func main() {
	conn, err := amqp.Dial("amqp://admin:123456@192.168.10.10:5672")
	if err != nil {
		fmt.Printf("connect to RabbitMQ failed, err:%v\n", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("Opena channel failed, err:%v\n", err)
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("declare a queue failed, err:%v\n", err)
		return
	}

	body := bodyFrom(os.Args)
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	if err != nil {
		fmt.Printf("publish a message failed, err:%v\n", err)
		return
	}
	log.Printf("[x] Sent %s", body)
}

// bodyFrom 从命令行获取将要发送的消息内容
func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
