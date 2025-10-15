package service

import (
	"github.com/rahulSailesh-shah/ch8n_go/internal/repo"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/config"
)

type Service struct {
	User UserService
}

func NewService(repos *repo.Repositories, cfg *config.AppConfig) *Service {
	return &Service{
		User: NewUserService(repos.User),
	}
}
