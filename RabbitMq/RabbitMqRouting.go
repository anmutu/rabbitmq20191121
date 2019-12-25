package RabbitMq

import (
	"github.com/streadway/amqp"
	"log"
)

//创建rabbitmq实例，这里有了routingkey为参数了。
func NewRabbitMqRouting(exchangeName string, routingKey string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchangeName, routingKey)
	var err error
	//获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)
	rabbitmq.failOnErr(err, "创建rabbit的路由实例的时候连接出现问题")
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "创建rabbitmq的路由实例时获取channel出错")
	return rabbitmq
}

//路由模式，产生消息。
func (r *RabbitMQ) PublishRouting(message string) {
	//第一步，尝试创建交换机，与pub/sub模式不同的是这里的kind需要是direct
	err := r.channel.ExchangeDeclare(r.ExChange, "direct", true, false, false, false, nil)
	r.failOnErr(err, "路由模式，尝试创建交换机失败")
	//第二步，发送消息
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

//路由模式，消费消息。
func (r *RabbitMQ) ConsumerRouting() {
	//第一步，尝试创建交换机，注意这里的交换机类型与发布订阅模式不同，这里的是direct
	err := r.channel.ExchangeDeclare(
		r.ExChange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "路由模式，创建交换机失败。")

	//第二步，尝试创建队列,注意这里队列名称不用写，这样就会随机产生队列名称
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "路由模式，创建队列失败。")

	//第三步，绑定队列到exchange中
	err = r.channel.QueueBind(q.Name, r.Key, r.ExChange, false, nil)

	//第四步，消费消息。
	messages, err := r.channel.Consume(q.Name, "", true, false, false, false, nil)
	forever := make(chan bool)
	go func() {
		for d := range messages {
			log.Printf("小杜同学写的路由模式收到消息为：%s。\n", d.Body)
		}
	}()
	<-forever

}
