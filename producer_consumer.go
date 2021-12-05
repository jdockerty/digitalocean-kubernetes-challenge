package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/segmentio/kafka-go"
)

const (
	topic     = "dok8s"
	partition = 0
	addr      = "my-cluster-kafka-bootstrap.kafka:9092"
)

// Our consumer continually reads from the topic, printing out messages from the producers.
func run_consumer() {

	log.Println("Starting consumer")

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{addr},
		Topic:     topic,
		Partition: partition,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		log.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}

}

// run_producer will send our defined messages into topic to be read by the consumer.
func run_producer() {
	log.Println("Running producer")

	w := &kafka.Writer{
		Addr:     kafka.TCP(addr),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Hello"),
			Value: []byte("World!"),
		},
		kafka.Message{
			Key:   []byte("Reason"),
			Value: []byte("Digital Ocean Kubernetes Challenge!"),
		},
		kafka.Message{
			Key:   []byte("Task"),
			Value: []byte("Scalable Message Queue"),
		},
	)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	log.Println("Messages written!")

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

	log.Println("Finished producing messages!")
}

func main() {

	entry := make(map[string]func())

	entry["CONSUMER"] = run_consumer
	entry["PRODUCER"] = run_producer

	// Map particular run function to environment variable of 'CONSUMER' or 'PRODUCER'
	entry[strings.ToUpper(os.Getenv("DOK8S_ROLE"))]()
}
