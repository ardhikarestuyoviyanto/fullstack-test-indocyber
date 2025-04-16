package service

import (
	"auth_test/config"
	"auth_test/internal/domain"
	"auth_test/internal/entity"
)
type svcAuth struct {
	c    config.Config
	repo domain.AdapterAuthService
}

func (s *svcAuth) LoginUsers(email string, password string) (entity.Users, bool) {
	return s.repo.LoginUsers(email, password)
}

func NewServiceAuth(repo domain.AdapterAuthRepository, c config.Config) domain.AdapterAuthService {
	return &svcAuth{
		c:    c,
		repo: repo,
	}
}

