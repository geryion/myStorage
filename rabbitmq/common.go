package rabbitmq

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

/*
	确定exchange交换类型
*/

//发送数据到exchange
func (mq *RabbitMQ)RPublish(exchange string, body interface{}) {
	str, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	err = mq.Channel.Publish(exchange, "", false, false, amqp.Publishing{
		ReplyTo: mq.Name,
		Body: []byte(str),
	})
	if err != nil {
		panic(err)
	}
}

//关闭传输通道
func (mq *RabbitMQ)RClose()  {
	mq.Channel.Close()
}

//绑定exchange
func (mq *RabbitMQ)RBind(exchange string) {
	err := mq.Channel.QueueBind(mq.Name, "", exchange, false, nil)
	if err != nil {
		panic(err)
	}
	mq.Exchange = exchange
}

//发送数据到queue
func (mq *RabbitMQ)RSend(queue string, body interface{})  {
	str, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	//均使用Publish进行数据的传输
	err = mq.Channel.Publish("", queue, false, false, amqp.Publishing{
		ReplyTo: mq.Name,
		Body: []byte(str),
	})
	if err != nil {
		panic(err)
	}
}

//接收消息channel
func (mq *RabbitMQ)RConsume() <-chan amqp.Delivery {
	ch, err := mq.Channel.Consume(mq.Name, "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	return ch
}