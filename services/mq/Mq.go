package mq

import (
	"github.com/streadway/amqp"
)

type Callback func(msg string)

func Connect() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	return conn, err
}

//发送端函数
func Publish(exchange string, queueName string, body string) error {
	//建立连接
	conn, err := Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	//创建通道channel
	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	//创建队列
	q, err := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	//发送消息
	err = channel.Publish(exchange, q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(body),
	})
	return err
}

func PublishDlx(exchangeA string, body string) error {
	//建立连接
	conn, err := Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	//创建一个Channel
	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	//消息发送到A交换机
	err = channel.Publish(exchangeA, "", false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(body),
	})

	return err
}
