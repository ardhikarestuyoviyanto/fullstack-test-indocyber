package domain

import "auth_test/internal/entity"

type AdapterAuthRepository interface {
	LoginUsers(email string, password string) (entity.Users, bool)
}

type AdapterAuthService interface {
	LoginUsers(email string, password string) (entity.Users, bool)
}