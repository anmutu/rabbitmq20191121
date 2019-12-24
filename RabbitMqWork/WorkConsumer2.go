package main

import "rabbitmq20181121/RabbitMq"

func main() {
	rabbitmq := RabbitMq.NewRabbitMQSimple("duQueueName191224")
	rabbitmq.ConsumeSimple()
}
