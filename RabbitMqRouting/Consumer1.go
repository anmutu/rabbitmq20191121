package main

import "rabbitmq20181121/RabbitMq"

func main() {
	one := RabbitMq.NewRabbitMqRouting("duExchangeName", "one")
	one.ConsumerRouting()
}
