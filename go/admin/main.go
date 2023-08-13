package main

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

var brokerAddrs = []string{"127.0.0.1:9095", "127.0.0.1:9096", "127.0.0.1:9097"}
var topicName = "web_log"

func main() {
	createTopic()
	time.Sleep(1 * time.Second)
	listTopic()
	describeTopic()
	deleteTopic()
}

func createTopic() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	admin, err := sarama.NewClusterAdmin(brokerAddrs, config)
	if err != nil {
		log.Fatal("Error while creating cluster admin: ", err.Error())
	}
	defer func() { _ = admin.Close() }()

	err = admin.CreateTopic(topicName, &sarama.TopicDetail{
		NumPartitions:     3,
		ReplicationFactor: 3,
	}, false)
	if err != nil {
		log.Fatal("Error while creating topic: ", err.Error())
	}
	fmt.Printf("create topic success : %s\n", topicName)
}

// 如果創建後馬上去查，沒那麼即時會顯示出來
func listTopic() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	admin, err := sarama.NewClusterAdmin(brokerAddrs, config)
	if err != nil {
		log.Fatal("Error while creating cluster admin: ", err.Error())
	}
	defer func() { _ = admin.Close() }()

	list, err := admin.ListTopics()
	if err != nil {
		log.Fatal("Error while list topic: ", err.Error())
	}
	fmt.Printf("list topic: %+v\n", list)
}

// 這個接口的開發貢獻者有點毛病，找不到topic的情況下
// admin.DescribeTopics 本身的error都沒回傳訊息，始終為nil
// topic.Err 會回傳 kafka server: Not an error, why are you printing me?
// 但我想能透過topic.Partition的長度來判斷該topic是否真實存在
func describeTopic() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	admin, err := sarama.NewClusterAdmin(brokerAddrs, config)
	if err != nil {
		log.Fatal("Error while creating cluster admin: ", err.Error())
	}
	defer func() { _ = admin.Close() }()

	list, err := admin.DescribeTopics([]string{topicName})
	if err != nil {
		log.Fatal("Error while list topic: ", err.Error())
	}
	fmt.Printf("describe topic: %+v\n", list)
	for _, topic := range list {
		fmt.Printf("topic name: %s, IsInternal: %t,partition_len: %d, version: %d,err: %v\n", topic.Name, topic.IsInternal, len(topic.Partitions), topic.Version, topic.Err)
	}

}

func deleteTopic() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	admin, err := sarama.NewClusterAdmin(brokerAddrs, config)
	if err != nil {
		log.Fatal("Error while creating cluster admin: ", err.Error())
	}
	defer func() { _ = admin.Close() }()

	err = admin.DeleteTopic(topicName)
	if err != nil {
		log.Fatal("Error while delete topic: ", err.Error())
	}
	fmt.Printf("delete topic success : %s\n", topicName)
}
