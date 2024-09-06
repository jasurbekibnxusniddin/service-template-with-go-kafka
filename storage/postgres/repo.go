package postgres

import (
	"github.com/jasurbekibnxusniddin/service-template-with-go-kafka/dto"
)

type TodoRepoI interface {
	Create(todo *dto.CreateTodoDto) error
}
