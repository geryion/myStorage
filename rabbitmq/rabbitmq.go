package rabbitmq

import (
	"github.com/streadway/amqp"
)

//实例化一个RabbitMQ对象处理当前的数据服务节点
func RNew(s string) *RabbitMQ {
	var ramq *RabbitMQ
	var err error

	//创建mq连接
	conn, err := amqp.Dial(s)
	if err != nil {
		panic(err)
	}

	//创建管道
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	//声明Queue
	q, err := ch.QueueDeclare("", false, true, false, false, nil)
	if err != nil {
		panic(err)
	}

	ramq.Channel = ch
	ramq.Name = q.Name

	return ramq
}
