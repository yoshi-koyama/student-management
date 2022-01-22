package store

import (
	"database/sql"
	db "student-management/db/sqlc"
	"student-management/model"
)

type SqlStore struct {
	*db.Queries
}

func (s SqlStore) CreateStudent(student model.Student) error {
	panic("implement me")
}

func (s SqlStore) GetStudents(gender string) []model.Student {
	panic("implement me")
}

func (s SqlStore) GetDetailStudent(id int) model.Student {
	panic("implement me")
}

func NewSqlStore(database *sql.DB) Store {
	return &SqlStore{Queries: db.New(database)}
}
