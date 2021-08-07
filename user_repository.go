package main

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type IUserRepository interface {
	GetUserByEmail(email string) (*User, error)
}

type UserRepository struct {
	users []*User
}

func NewUserRepository() IUserRepository {
	p1, _ := bcrypt.GenerateFromPassword([]byte("11111111"), bcrypt.DefaultCost)
	p2, _ := bcrypt.GenerateFromPassword([]byte("22222222"), bcrypt.DefaultCost)

	users := []*User{
		&User{
			ID: 1,
			Email: "alex@example.com",
			Password: string(p1),
		},
		&User{
			ID: 2,
			Email: "mary@example.com",
			Password: string(p2),
		},
	}

	return &UserRepository{users: users}
}

func (r *UserRepository) GetUserByEmail(email string) (*User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}
