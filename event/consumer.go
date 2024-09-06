package event

import (
	"log"

	"github.com/jasurbekibnxusniddin/service-template-with-go-kafka/storage"
)

type event struct {
	storage storage.StorageI
}

func (e *event) createTodo() {

	topic := "create_todo"
	log.Println(topic, "consumer is running")
	consumeMessages(topic, e.storage)

}

func (e *event) updateTodo() {
	topic := "update_todo"
	log.Println(topic, "consumer is running")
	consumeMessages(topic, e.storage)
}

func (e *event) Run() {
	go e.createTodo()
	go e.updateTodo()
	select {}
}

func NewEvent(storage storage.StorageI) *event {
	return &event{storage: storage}
}
