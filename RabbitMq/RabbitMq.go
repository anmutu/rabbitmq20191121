package RabbitMq

import (
	"fmt"
	"github.com/streadway/amqp"
)

//连接信息
const MQURL = "amqp://du:du@129.211.78.6:5672/dudevirtualhost"

//RabbitMQ结构体
type RabbitMQ struct {
	//连接
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列
	QueueName string
	//交换机名称
	ExChange string
	//绑定的key名称
	Key string
	//连接的信息，上面已经定义好了
	MqUrl string
}

//创建结构体实例，参数队列名称、交换机名称和bind的key（也就是几个大写的，除去定义好的常量信息）
func NewRabbitMQ(queueName string, exChange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, ExChange: exChange, Key: key, MqUrl: MQURL}
}

//关闭conn和chanel的方法
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

//错误的函数处理
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		fmt.Printf("err是:%s,小杜同学手写的信息是:%s", err, message)
	}
}
