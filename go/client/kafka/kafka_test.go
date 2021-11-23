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
	partitionList, _ := GetConsumer().Partitions("topic")
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
