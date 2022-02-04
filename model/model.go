package model

type Student struct {
	Id       int64  `json:"id"`
	Name     string `json:"name" binding:"required,min=3,max=16"`
	Nrc      string `json:"nrc" binding:"required"` //unique
	Age      int32  `json:"age" binding:"required"`
	Gender   string `json:"gender" binding:"required,oneof=male female"`
	Password string `json:"password" binding:"required,min=7,max=16"`
}

type GetDetailStudentRequest struct {
	Id int64 `uri:"id" binding:"required"`
}

type GetFilteredStudentRequest struct {
	Gender string `form:"gender" binding:"omitempty,oneof=male female"`
	Nrc string `form:"nrc"`
	Age int32 `form:"age"`
}

type AuthHeader struct {
	Token string `header:"token" binding:"required"`
}
