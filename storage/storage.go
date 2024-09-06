package storage

import (
	"github.com/jasurbekibnxusniddin/service-template-with-go-kafka/storage/postgres"
	"gorm.io/gorm"
)

type StorageI interface {
	GetTodoRepo() postgres.TodoRepoI
}

type storage struct {
	todoRrepo postgres.TodoRepoI
}

func NewStorage(db *gorm.DB) StorageI {
	return &storage{
		todoRrepo: postgres.NewTodoRepo(db),
	}
}

func (s *storage) GetTodoRepo() postgres.TodoRepoI {
	return s.todoRrepo
}
