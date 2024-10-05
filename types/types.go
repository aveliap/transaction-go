package types

import "time"

type UserRepo interface{
	GetUserByEmail (email string) (*User, error)
	GetUserByID (id uint) (*User, error)
	CreateUser (User) error
}

type User struct{
	ID uint `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayload struct{
	FirstName string `json:"firstName" validate:"required"`
	LastName string `json:"lastName" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=8"`

}

type loginUserPayload struct{
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`

}