package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

var rabbitmqCh *amqp.Channel

func InitAmqp() {
	conn, err := amqp.Dial("amqp://admin:123456@localhost:5672/")
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ:", err)
	}
	defer conn.Close()

	// 创建通道
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Failed to open a channel:", err)
	}
	defer ch.Close()

	// 申明队列
	q, err := ch.QueueDeclare(
		"hello", // 队列名称
		false,   // 是否持久化
		false,   // 是否自动删除
		false,   // 是否排他性
		false,   // 是否等待服务器响应
		nil,     // 额外的属性
	)
	if err != nil {
		fmt.Println("Failed to declare a queue:", err)
	}

	// 向队列发送消息
	body := "Hello World!"
	err = ch.Publish(
		"",     // 交换机名称
		q.Name, // 路由键
		false,  // 是否强制发送到队列
		false,  // 是否等待服务器响应
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		fmt.Println("Failed to publish a message:", err)
	}
	// 从队列接受消息
	msgs, err := ch.Consume(
		q.Name, // 队列名称
		"",     // 消费者名称
		true,   // 是否自动响应确认
		false,  // 是否排他性
		false,  // 是否阻塞等待新消息
		false,  // 是否等待服务器响应
		nil,    // 额外的属性
	)
	if err != nil {
		fmt.Println("Failed to register a consumer:", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Println("Received a message", string(d.Body))
		}
	}()

	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
