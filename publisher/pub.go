package pub

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func Pub() {
	ctx := context.Background()
	proj := "root-furnace-345606"

	client, err := pubsub.NewClient(ctx, proj, option.WithCredentialsFile("root-furnace-345606-0daf3c5fb54d.json"))
	if err != nil {
		log.Fatalf("Could not create pubsub Client: %v", err)
	}
	const topic = "payments"
	// Create a new topic called my-topic.
	// if err := pub.Create(client, topic); err != nil {
	// 	log.Fatalf("Failed to create a topic: %v", err)
	// }
	// list all the topics from the project.
	fmt.Println("Listing all topics from the project:")
	topics, err := list(client)
	if err != nil {
		log.Fatalf("Failed to list topics: %v", err)
	}
	for _, t := range topics {
		fmt.Println(t)
	}
	// Publish a text message on the created topic.
	if err := publish(client, topic, "hello world!"); err != nil {
		log.Fatalf("Failed to publish: %v", err)
	}
}

func list(client *pubsub.Client) ([]*pubsub.Topic, error) {
	ctx := context.Background()
	var topics []*pubsub.Topic
	it := client.Topics(ctx)
	for {
		topic, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		topics = append(topics, topic)
	}
	return topics, nil
}
func Create(client *pubsub.Client, topic string) error {
	ctx := context.Background()
	t, err := client.CreateTopic(ctx, topic)
	if err != nil {
		return err
	}
	fmt.Printf("Topic created: %v\n", t)
	return nil
}
func publish(client *pubsub.Client, topic, msg string) error {
	ctx := context.Background()
	t := client.Topic(topic)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
	return nil
}
