package event

import (
	"context"
	"encoding/json"
	"log"

	"github.com/jasurbekibnxusniddin/service-template-with-go-kafka/dto"
	"github.com/jasurbekibnxusniddin/service-template-with-go-kafka/storage"
	"github.com/segmentio/kafka-go"
)

const (
	KafkaBroker = "localhost:9092"
	GroupID     = "crud-service-group"
)

func consumeMessages(topic string, storage storage.StorageI) {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{KafkaBroker},
		Topic:   topic,
		//GroupID: GroupID,
	})

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("could not read message:", err)
			continue
		}

		switch topic {
		case "create_todo":
			var todo dto.CreateTodoDto
			if err := json.Unmarshal(msg.Value, &todo); err != nil {
				log.Println("could not unmarshal message:", err)
				continue
			}

			err := storage.GetTodoRepo().Create(&todo)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
}
