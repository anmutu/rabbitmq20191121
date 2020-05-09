#### 这里用golang实现了rabbigmq的一种模式
1. simple模式
2. work模式
3. 订阅模式
4. routing模式
5. topic模式

#### 文件介绍
1. RabbitMq文件

1.1 RabbitMq文件夹里的RabbitMq.go文件是其结构体和其函数。

1.2 RabbitMq文件夹里其他go文件则是rabbitmq的各种模式的代码逻辑了，其中simple模式和     worker模式都是一样的go文件。

2. 除去RabbitMq的其他文件夹为各种模式的生产者和消费者的文件了。

 #### 如何run起来？
 每个模式里都有其对应的生产者和消费者，分别run起来里面的生产者的go文件和消费者的go文件就可以看到效果了。
 如果你有自己的rabbitmq也可以改成自己的ip和对应的exchangeName等信息。

   

   

   










