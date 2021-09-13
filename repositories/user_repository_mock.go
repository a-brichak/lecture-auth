package repositories

import (
	"auth/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryMock struct {
	users []*models.User
}

func NewUserRepositoryMock() IUserRepository {
	p1, _ := bcrypt.GenerateFromPassword([]byte("test_passw"), bcrypt.DefaultCost)
	p2, _ := bcrypt.GenerateFromPassword([]byte("test_passw_2"), bcrypt.DefaultCost)

	users := []*models.User{
		&models.User{
			ID:       1,
			Email:    "test-1@example.com",
			Name:     "Test User 1",
			Password: string(p1),
		},
		&models.User{
			ID:       2,
			Email:    "test-2@example.com",
			Name:     "TestUser2",
			Password: string(p2),
		},
	}

	return &UserRepositoryMock{users: users}
}

func (r *UserRepositoryMock) GetUserByEmail(email string) (*models.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (r *UserRepositoryMock) GetUserByID(id int) (*models.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}
