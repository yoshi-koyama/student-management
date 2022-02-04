package store

import (
	"context"
	"database/sql"
	db "student-management/db/sqlc"
	"student-management/model"
)

type SqlStore struct {
	*db.Queries
}

func (s SqlStore) CreateStudent(student model.Student) error {
	p := db.CreateStudentParams{Name: student.Name, Nrc: student.Nrc, Age: student.Age, Gender: student.Gender, Password: student.Password}
	return s.Queries.CreateStudent(context.Background(), p)
}

func (s SqlStore) GetStudents(gender string) ([]model.Student, error) {
	students, err := s.Queries.ListStudents(context.Background())
	if err != nil {
		return nil, err
	}
	modelStudents := make([]model.Student, len(students))
	for i, v := range students {
		modelStudents[i] = model.Student{Id: v.ID, Name: v.Name, Nrc: v.Nrc, Age: v.Age, Gender: v.Gender, Password: v.Password}
	}
	return modelStudents, nil
}

func (s SqlStore) GetDetailStudent(id int64) (*model.Student, error) {
	student, err := s.Queries.FindById(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &model.Student{Id: student.ID, Name: student.Name, Nrc: student.Nrc, Age: student.Age, Gender: student.Gender, Password: student.Password}, nil
}

func NewSqlStore(database *sql.DB) Store {
	return &SqlStore{Queries: db.New(database)}
}
