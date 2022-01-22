package store

import "student-management/model"

type Store interface {
	CreateStudent(student model.Student) error
	GetStudents(gender string) ([]model.Student, error)
	GetDetailStudent(id int64) (*model.Student, error)
}
