package main

import "rabbitmq20181121/RabbitMq"

func main() {
	two := RabbitMq.NewRabbitMqRouting("duExchangeName", "two")
	two.ConsumerRouting()
}
