package RabbitMq

import (
	"github.com/streadway/amqp"
	"log"
)

//topic模式
//与routing模式不同的是这个exchange的kind是"topic"类型的。
//topic模式的特别是可以以通配符的形式来指定与之匹配的消费者。
//"*"表示匹配一个单词。“#”表示匹配多个单词，亦可以是0个。

//创建rabbitmq实例
func NewRabbitMqTopic(exchangeName string, routingKey string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchangeName, routingKey)
	var err error
	//获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)
	rabbitmq.failOnErr(err, "创建rabbit的topic模式时候连接出现问题")
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "创建rabbitmq的topic实例时获取channel出错")
	return rabbitmq
}

//topic模式。生产者。
func (r *RabbitMQ) PublishTopic(message string) {
	//第一步，尝试创建交换机,这里的kind的类型要改为topic
	err := r.channel.ExchangeDeclare(
		r.ExChange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "topic模式尝试创建exchange失败。")

	//第二步，发送消息。
	err = r.channel.Publish(
		r.ExChange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

//topic模式。消费者。"*"表示匹配一个单词。“#”表示匹配多个单词，亦可以是0个。
func (r *RabbitMQ) ConsumerTopic() {
	//第一步，创建交换机。这里的kind需要是“topic”类型。
	err := r.channel.ExchangeDeclare(
		r.ExChange,
		"topic",
		true, //这里需要是true
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "topic模式，消费者创建exchange失败。")

	//第二步，创建队列。这里不用写队列名称。
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "topic模式，消费者创建queue失败。")

	//第三步，将队列绑定到交换机里。
	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.ExChange,
		false,
		nil,
	)

	//第四步，消费消息。
	messages, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range messages {
			log.Printf("小杜同学写的topic模式收到了消息：%s。\n", d.Body)
		}
	}()
	<-forever

}
