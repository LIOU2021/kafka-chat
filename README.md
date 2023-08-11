# Ref
- 觀念
    - [Kafka 介紹 + Golang程式實作](https://ftn8205.medium.com/kafka-%E4%BB%8B%E7%B4%B9-golang%E7%A8%8B%E5%BC%8F%E5%AF%A6%E4%BD%9C-2b108481369e)
- 實作
    - [docker 配置 kafka+zookeeper，golang操作kafka](https://blog.51cto.com/u_6192297/3299886)
    - [Kafka 生产者和消费者学习笔记](https://leehao.me/Kafka-%E7%94%9F%E4%BA%A7%E8%80%85%E5%92%8C%E6%B6%88%E8%B4%B9%E8%80%85%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0/)

# tip
- 我是在windows上測試的
- ./go/consumer_group的這個環境沒法在windows模擬，因為win沒有SIGUSR1訊號，所以我是透過docker在容器內執行測試這個容器的