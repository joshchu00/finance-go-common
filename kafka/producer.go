package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	producer *kafka.Producer
}

func NewProducer(servers string) (p *Producer, err error) {

	p = &Producer{}

	if p.producer, err = kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers": servers,
			// "linger.ms":          "50",
			// "batch.num.messages": "1000",
			// "batch.size":"100000",
			"default.topic.config": kafka.ConfigMap{
				"request.required.acks": "all",
				"message.timeout.ms":    "0",
			},
		},
	); err != nil {
		return
	}

	return
}

func (p *Producer) Produce(topic string, partition int32, value []byte) (err error) {

	if err = p.producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: partition,
			},
			Value: value,
		},
		nil,
	); err != nil {
		return
	}

	return
}

func (p *Producer) Flush(timeoutMs int) {
	p.producer.Flush(timeoutMs)
}

// Close
func (p *Producer) Close() {
	p.producer.Close()
}
