package service

import (
	"context"
	"fmt"

	"github.com/rahulSailesh-shah/ch8n_go/internal/dto"
	"github.com/rahulSailesh-shah/ch8n_go/internal/model"
	"github.com/rahulSailesh-shah/ch8n_go/internal/repo"
)

var (
	ErrUserNotFound      = fmt.Errorf("user not found")
	ErrUsernameExists    = fmt.Errorf("username already exists")
	InvalidCredentials   = fmt.Errorf("invalid credentials")
	ErrUserAlreadyExists = fmt.Errorf("user with this username already exists")
)

type UserService interface {
	CreateUser(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error)
}

type userService struct {
	repo repo.UserRepo
}

func NewUserService(repo repo.UserRepo) UserService {
	return &userService{
		repo: repo,
	}
}

func (u *userService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error) {
	existingUser, err := u.repo.GetUserByEmail(ctx, req.Email)
	if err == nil && existingUser != nil {
		return nil, ErrUserAlreadyExists
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	newUser, err := u.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return toUserResponse(newUser), nil
}

func toUserResponse(user *model.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
