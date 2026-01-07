package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/FitRang/fanout-consumer/eventbus"
	"github.com/FitRang/fanout-consumer/redispub"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	cfg := eventbus.Config{
		Brokers: os.Getenv("KAFKA_URI"),
	}

	bus, err := eventbus.NewEventBus(cfg)
	if err != nil {
		log.Fatalf("failed to init event bus: %v", err)
	}

	rdb, err := redispub.NewRedisClient(os.Getenv("REDIS_URI"))
	if err != nil {
		log.Fatal("while creating redis clien:", err)
	}
	log.Printf("redis connected:%s", os.Getenv("REDIS_URI"))

	consumer, err := bus.NewConsumer(
		"fanout-consumer",
		[]string{"notification"},
		func(key, value []byte) error {
			if err := redispub.PublishToRedis(rdb, value); err != nil {
				log.Printf("redis publish failed: %v", err)
			}
			return nil
		},
	)
	if err != nil {
		log.Fatalf("failed to create consumer: %v", err)
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	log.Println("Kafka consumer started")

	<-ctx.Done()

	log.Println("shutting down consumer...")
	consumer.Close()

	log.Println("consumer stopped cleanly")
}
