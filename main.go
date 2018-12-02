package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// ExampleBroker to get metadata
func ExampleBroker() {
	broker := sarama.NewBroker("192.168.99.102:9092")
	err := broker.Open(nil)
	if err != nil {
		panic(err)
	}

	request := sarama.MetadataRequest{Topics: []string{"test", "topic", "something", "cool"}}
	response, err := broker.GetMetadata(&request)
	if err != nil {
		_ = broker.Close()
		panic(err)
	}

	client, err := sarama.NewClient([]string{broker.Addr()}, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Partitions: ")
	fmt.Println(client.Topics())

	fmt.Println("There are", len(response.Topics), "topics active in the cluster.")

	fmt.Println(response.ClusterID)

	if err = broker.Close(); err != nil {
		panic(err)
	}
}

func main() {

	ExampleBroker()

}
