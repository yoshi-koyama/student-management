package store

import "student-management/model"

type Store interface {
	CreateStudent(student model.Student) error
	GetStudents(gender string) []model.Student
	GetDetailStudent(id int) model.Student
}
