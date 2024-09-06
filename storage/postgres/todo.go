package postgres

import (
	"github.com/google/uuid"
	"github.com/jasurbekibnxusniddin/service-template-with-go-kafka/dto"
	"github.com/jasurbekibnxusniddin/service-template-with-go-kafka/models"
	"gorm.io/gorm"
)

type todoRepo struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) TodoRepoI {
	return &todoRepo{db: db}
}

func (t *todoRepo) Create(todoDto *dto.CreateTodoDto) error {

	todo := models.TodoModel{
		Id:   uuid.New(),
		Task: todoDto.Task,
	}

	tx := t.db.Create(&todo)

	return tx.Error
}
