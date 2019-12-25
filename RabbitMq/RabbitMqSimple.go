package RabbitMq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//创建简单模式下的实例，只需要queueName这个参数，其中exchange是默认的，key则不需要。
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	rabbitmq := NewRabbitMQ(queueName, "", "")
	var err error
	//获取参数connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)
	rabbitmq.failOnErr(err, "连接connection失败")
	//获取channel参数
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel参数失败")
	return rabbitmq
}

//直接模式,生产者.
func (r *RabbitMQ) PublishSimple(message string) {
	//第一步，申请队列，如不存在，则自动创建之，存在，则路过。
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("创建连接队列失败：%s", err)
	}

	//第二步，发送消息到队列中
	r.channel.Publish(
		r.ExChange,
		r.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

//直接模式，消费者
func (r *RabbitMQ) ConsumeSimple() {
	//第一步,申请队列,如果队列不存在则自动创建,存在则跳过
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	//第二步,接收消息
	msgs, err := r.channel.Consume(
		q.Name,
		"",   //用来区分多个消费者
		true, //是否自动应答,告诉我已经消费完了
		false,
		false, //若设置为true,则表示为不能将同一个connection中发送的消息传递给这个connection中的消费者.
		false, //消费队列是否设计阻塞
		nil,
	)
	if err != nil {
		fmt.Printf("消费者接收消息出现问题:%s", err)
	}

	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for d := range msgs {
			log.Printf("小杜同学写的Simple模式接收到了消息:%s。\n", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
