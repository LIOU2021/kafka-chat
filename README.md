# Ref
- 觀念
    - [Kafka 介紹 + Golang程式實作](https://ftn8205.medium.com/kafka-%E4%BB%8B%E7%B4%B9-golang%E7%A8%8B%E5%BC%8F%E5%AF%A6%E4%BD%9C-2b108481369e)
    - [【总结】kafka-topics.sh --describe显示结果解释](https://developer.aliyun.com/article/515351)
- 實作
    - [docker 配置 kafka+zookeeper，golang操作kafka](https://blog.51cto.com/u_6192297/3299886)
    - [Kafka 生产者和消费者学习笔记](https://leehao.me/Kafka-%E7%94%9F%E4%BA%A7%E8%80%85%E5%92%8C%E6%B6%88%E8%B4%B9%E8%80%85%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0/)
    - [kafka笔记](https://juejin.cn/post/6844903887149760526)

# tip
- 我是在windows上測試的
- ./go/consumer_group的這個環境沒法在windows模擬，因為win沒有SIGUSR1訊號，所以我是透過docker在容器內執行測試這個容器的

# kafka cli
```sh
# 查看Partition數量
$KAFKA_HOME/bin/kafka-topics.sh --describe --zookeeper zookeeper:2181 --topic web_log
# 修改Partition數量為3
$KAFKA_HOME/bin/kafka-topics.sh --zookeeper zookeeper:2181 --alter --topic web_log --partitions 3
# 檢視有多少broker數量
$KAFKA_HOME/bin/zookeeper-shell.sh zookeeper:2181 ls /brokers/ids

# 創建一個topic名為web_log，partition為3，replica為3
$KAFKA_HOME/bin/kafka-topics.sh --create --topic web_log --partitions 3 --zookeeper zookeeper:2181 --replication-factor 3
# 查看上述指令創建結果
$KAFKA_HOME/bin/kafka-topics.sh --describe --zookeeper zookeeper:2181 --topic web_log
# 列出目前所有的topic
$KAFKA_HOME/bin/kafka-topics.sh --list --zookeeper zookeeper:2181
# 刪除topic
$KAFKA_HOME/bin/kafka-topics.sh --delete --topic web_log --zookeeper zookeeper:2181
# 生產者傳送訊息
$KAFKA_HOME/bin/kafka-console-producer.sh --topic web_log --bootstrap-server kafka1:9094,kafka2:9093,kafka3:9094
# 消費者接收訊息
$KAFKA_HOME/bin/kafka-console-consumer.sh --topic web_log --from-beginning --bootstrap-server host.docker.internal:9095,host.docker.internal:9096,host.docker.internal:9097
```

# 白話文重點
- 一個topic底下可以有多個Partition，這個partition就像是queue，存放著訊息
- client要訂閱的話，可以直接以消費者身分訂閱partition，多個client以此方式訂閱的話，就像是redis channel的訂閱依樣，所有客戶端都會收到一樣的訊息
- consumer訂閱的目標其實是partition，單一consumer可以訂閱多個partition達成完整接收特定topic訊息的目的
- client還可以用consumer group 的身分基礎去訂閱，如此一來群組成員的接收訊息就會不一致了，但該consumer group底下的所有consumer 收到的訊息彙總，就是完整的訂閱訊息了
- 如果partition只有1個，那麼採用consumer group的情況下，該group如果有兩個(包含)以上的成員，就永遠只會有特定人會收到訊息，另一個人則接收不到訊息
- 如果partition有1個以上，consumer group底下不同成員間都會收到片段訊息。因為每個成員都是訂閱不同的partition的概念。不同的consumer group可以訂閱相同的topic或是partition，但同一consumer group底下的成員不能訂閱重複相同的partition
- broker其實就是kafka架設了幾個服務