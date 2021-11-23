package kafka

import (
	"time"

	"github.com/Shopify/sarama"
)

var syncProducer sarama.SyncProducer
var asyncProducer sarama.AsyncProducer
var consumer sarama.Consumer

var eventSyncProducer sarama.SyncProducer
var eventAsyncProducer sarama.AsyncProducer
var eventConsumer sarama.Consumer

func InitKafka() {
	config := sarama.NewConfig()
	// Producer
	// 重试
	config.Producer.Retry.Max = 5
	// 1s
	config.Producer.Retry.Backoff = 100 * time.Second
	config.Producer.Retry.BackoffFunc = nil
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	// v0_10_0_0开始支持消息Timestamp
	version, err := sarama.ParseKafkaVersion("0.10.0.0")
	if err != nil {
		panic(err)
	}
	config.Version = version

	// syncProducer
	config.Producer.RequiredAcks = sarama.WaitForAll // 消息可靠性，等待所有副本提交成功

	syncProducer, err = sarama.NewSyncProducer([]string{"127.0.0.1:2181"}, config)
	if err != nil {
		panic(err)
	}

	// asyncProducer
	config.Producer.RequiredAcks = sarama.WaitForLocal       // 消息可靠性，只关心leader提交成功
	config.Producer.Compression = sarama.CompressionSnappy   // 消息压缩
	config.Producer.Flush.Frequency = 500 * time.Millisecond // 消息发送频率

	asyncProducer, err = sarama.NewAsyncProducer([]string{"127.0.0.1:2181"}, config)
	if err != nil {
		panic(err)
	}

	consumer, err = sarama.NewConsumer([]string{"127.0.0.1:2181"}, config)
	if err != nil {
		panic(err)
	}
}

func GetConsumer() sarama.Consumer {
	return consumer
}
