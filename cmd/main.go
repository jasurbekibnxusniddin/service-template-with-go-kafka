package main

import (
	"fmt"

	"github.com/jasurbekibnxusniddin/service-template-with-go-kafka/event"
	"github.com/jasurbekibnxusniddin/service-template-with-go-kafka/storage"
	"github.com/jasurbekibnxusniddin/service-template-with-go-kafka/storage/postgres"
)

func main() {

	db, err := postgres.ConnToDB()
	if err != nil {
		return
	}

	err = postgres.AoutoMigrate(db)
	if err != nil {
		return
	}

	storage := storage.NewStorage(db)

	event := event.NewEvent(storage)

	event.Run()

	fmt.Println(err)
}
