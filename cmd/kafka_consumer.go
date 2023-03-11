package cmd

import (
	"database/sql"
	"encoding/json"
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/mattn/go-sqlite3"
	"go-notification-api/internal/infrastructure/database"
	"go-notification-api/internal/usecase"
	"go-notification-api/pkg/kafka"
	"log"
)

func InitConsumer(db *sql.DB) {
	repository := database.NewNotificationRepository(db)
	useCase := usecase.NewNotificationFromKafka{NotificationRepository: repository}

	msgChanKafka := make(chan *ckafka.Message)
	topics := []string{"notifications"}
	servers := "host.docker.internal:9094"
	fmt.Println("Kafka consumer has started")
	go kafka.Consume(topics, servers, msgChanKafka)
	worker(msgChanKafka, useCase)
}

func worker(msgChan chan *ckafka.Message, uc usecase.NewNotificationFromKafka) {
	fmt.Println("Kafka worker has started")
	for msg := range msgChan {
		var notificationInputDTO usecase.NotificationInputDTO
		err := json.Unmarshal(msg.Value, &notificationInputDTO)
		if err != nil {
			panic(err)
		}
		notificationOutputDto, err := uc.Execute(notificationInputDTO)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("Kafka has processed notification %d\n", notificationOutputDto.ID)
	}
}
