package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

// kafka consumer

var Topic = "web_log" //主题名称

func main() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9095", "127.0.0.1:9096", "127.0.0.1:9097"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions(Topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(Topic, int32(partition), sarama.OffsetNewest) // 新消費者上線後只會讀取最新訊息
		// pc, err := consumer.ConsumePartition(Topic, int32(partition), sarama.OffsetOldest) // 新消費者上線後會從第一則訊息開始讀
		// pc, err := consumer.ConsumePartition(Topic, int32(partition), 2) // 指定從offset 2 開始讀，要注意是不是該partition下有這個位置
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
			}
		}(pc)
	}
	select {} //阻塞进程
}
