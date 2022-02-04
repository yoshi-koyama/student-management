package db

import "student-management/model"

func (s Student) ToResponse() model.Student{
	return model.Student{
		Id:       s.ID,
		Name:     s.Name,
		Nrc:      s.Nrc,
		Age:      s.Age,
		Gender:   s.Gender,
		Password: "",
	}
}