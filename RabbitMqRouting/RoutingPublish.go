package main

import (
	"fmt"
	"rabbitmq20181121/RabbitMq"
	"strconv"
	"time"
)

func main() {
	rabbitmq1 := RabbitMq.NewRabbitMqRouting("duExchangeName", "one")
	rabbitmq2 := RabbitMq.NewRabbitMqRouting("duExchangeName", "two")
	rabbitmq3 := RabbitMq.NewRabbitMqRouting("duExchangeName", "three")
	for i := 0; i < 100; i++ {
		rabbitmq1.PublishRouting("路由模式one" + strconv.Itoa(i))
		rabbitmq2.PublishRouting("路由模式two" + strconv.Itoa(i))
		rabbitmq3.PublishRouting("路由模式three" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Printf("在路由模式下，routingKey为one,为two,为three的都分别生产了%d条消息\n", i)
	}
}
