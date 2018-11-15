package kafka

import (
	"log"
	"testing"
)

func TestProducer(t *testing.T) {

	var p *Producer
	var err error

	if p, err = NewProducer("192.168.33.10:9092"); err != nil {
		log.Panicln(err)
	}
	defer p.Close()

	if err = p.Produce("test", 0, []byte("{}")); err != nil {
		log.Panicln(err)
	}

	p.Flush(1000)
}
func TestConsumer(t *testing.T) {

	var c *Consumer
	var err error

	if c, err = NewConsumer("192.168.33.10:9092", "id1", "test"); err != nil {
		log.Panicln(err)
	}
	defer c.Close()

	var topic string
	var partition int32
	var offset int64
	var value []byte

	for {
		if topic, partition, offset, value, err = c.Consume(); err != nil {
			log.Panicln(err)
		}

		log.Println(topic, partition, offset, value)

		// if err = c.Commit(); err != nil {
		// 	log.Panicln(err)
		// }

		// strange
		offset++

		if err = c.CommitOffset(topic, partition, offset); err != nil {
			log.Panicln(err)
		}
	}
}
