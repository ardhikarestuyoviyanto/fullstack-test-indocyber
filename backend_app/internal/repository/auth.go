package repository

import (
	"auth_test/internal/domain"
	"auth_test/internal/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type repositoryAuth struct {
	DB *gorm.DB
}

func (r *repositoryAuth) LoginUsers(email string, password string) (entity.Users, bool) {
	users := entity.Users{}
	err_email := r.DB.Where("email = ?", email).First(&users).Error
	if err_email != nil {
		return users, false
	}
	err_password := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))
	if err_password != nil && err_password == bcrypt.ErrMismatchedHashAndPassword {
		return users, false
	}
	return users, true
}


func NewAuthRepository(db *gorm.DB) domain.AdapterAuthRepository {
	return &repositoryAuth{
		DB: db,
	}
}