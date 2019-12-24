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
	for i := 0; i < 100; i++ {
		rabbitmq1.PublishRouting("路由模式one" + strconv.Itoa(i))
		rabbitmq2.PublishRouting("路由模式two" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Printf("路由模式产生了%v条消息\n", i)
	}

}
