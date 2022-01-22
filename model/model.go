package model

type Student struct {
	Id       int    `json:"id"`
	Name     string `json:"name" binding:"required,min=3,max=16"`
	Nrc      string `json:"nrc" binding:"required"` //unique
	Gender   string `json:"gender" binding:"required,oneof=male female"`
	Password string `json:"password" binding:"required,min=7,max=16"`
}

type GetDetailStudentRequest struct {
	Id int `uri:"id" binding:"required"`
}

type GetFilteredStudentRequest struct {
	Gender string `form:"gender" binding:"omitempty,oneof=male female"`
}

type AuthHeader struct {
	Token string `header:"token" binding:"required"`
}
