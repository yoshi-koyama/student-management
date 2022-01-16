package store

import (
	"errors"
	"student-management/model"
)

type FakeStore struct {
	data map[string]model.Student
}

func NewStore() Store {
	return &FakeStore{data: make(map[string]model.Student)}
}

func (store *FakeStore) CreateStudent(student model.Student) error {
	_, ok := store.data[student.Nrc]
	if ok {
		return errors.New("nrc already exists")
	}

	student.Id = len(store.data) + 1

	store.data[student.Nrc] = student

	return nil
}
func (store *FakeStore) GetStudents(gender string) []model.Student {
	result := []model.Student{}
	for _, v := range store.data {
		if len(gender) > 0 {
			if v.Gender == gender {
				result = append(result, v)
			}
		} else {
			result = append(result, v)
		}
	}
	return result
}

func (store *FakeStore) GetDetailStudent(id int) model.Student {
	var student model.Student
	for _, v := range store.data {
		if id == v.Id {
			student = v
			break
		}
	}
	return student
}
