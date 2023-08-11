package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

var (
	Topic          = "web_log"
	ConsumerGroup1 = "my-consumer-group-1" // 修改为你的消费者群组 ID
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange // 可以根据需要选择负载均衡策略

	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(Topic, 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Printf("failed to start consumer for partition,err:%v\n", err)
		return
	}
	defer partitionConsumer.Close()

	fmt.Println("Consumer started")
	for msg := range partitionConsumer.Messages() {
		fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
	}
}
