package request

import "time"

type UserRequest struct {
	Email       string    `json:"email" binding:"required,email"`
	Password    string    `json:"password" binding:"required,min=6,containsany=!@#$%*"`
	Name        string    `json:"name" binding:"required,min=4,max=100"`
	Age         int8      `json:"age" binding:"required,min=1,max=140"`
	Nationality string    `json:"nationality" binding:"required,min=3,max=100"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserUpdateRequest struct {
	Name        string `json:"name" binding:"omitempty,min=4,max=100"`
	Age         int8   `json:"age" binding:"omitempty,min=1,max=140"`
	Nationality string `json:"nationality" binding:"omitempty,min=3,max=100"`
}
