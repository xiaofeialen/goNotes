package kafka

import (
	"fmt"
	"testing"

	"github.com/Shopify/sarama"
)

func init() {
	InitKafka()

}

func TestGetKafkaAsyncProducer(t *testing.T) {
	consumer, _ := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	partitionList, _ := consumer.Partitions("topic")
	for partition := range partitionList {
		pc, err := GetConsumer().ConsumePartition("topic", int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		for msg := range pc.Messages() {
			fmt.Println(msg)
		}
	}

}
