package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	consumer *kafka.Consumer
}

func NewConsumer(servers string, groupID string, topic string) (c *Consumer, err error) {

	c = &Consumer{}

	if c.consumer, err = kafka.NewConsumer(
		&kafka.ConfigMap{
			"bootstrap.servers":  servers,
			"group.id":           groupID,
			"enable.auto.commit": "false",
		},
	); err != nil {
		return
	}

	if err = c.consumer.Subscribe(topic, nil); err != nil {
		return
	}

	return
}

func (c *Consumer) Consume() (topic string, partition int32, offset int64, value []byte, err error) {

	var message *kafka.Message
	message, err = c.consumer.ReadMessage(-1)
	if err != nil {
		return
	}

	topic = *message.TopicPartition.Topic
	partition = message.TopicPartition.Partition
	offset = int64(message.TopicPartition.Offset)
	value = message.Value

	return
}

func (c *Consumer) Commit() (err error) {

	if _, err = c.consumer.Commit(); err != nil {
		return
	}

	return
}

func (c *Consumer) CommitOffset(topic string, partition int32, offset int64) (err error) {

	if _, err = c.consumer.CommitOffsets(
		[]kafka.TopicPartition{
			kafka.TopicPartition{
				Topic:     &topic,
				Partition: partition,
				Offset:    kafka.Offset(offset),
			},
		},
	); err != nil {
		return
	}

	return
}

func (c *Consumer) Close() {
	c.consumer.Close()
}
