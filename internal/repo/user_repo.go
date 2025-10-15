package repo

import (
	"context"

	"github.com/rahulSailesh-shah/ch8n_go/internal/db/sqlc"
	"github.com/rahulSailesh-shah/ch8n_go/internal/model"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userRepo struct {
	queries *sqlc.Queries
}

func NewUserRepo(queries *sqlc.Queries) UserRepo {
	return &userRepo{
		queries: queries,
	}
}

func (r *userRepo) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	newUser, err := r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return nil, err
	}
	return toDomain(&newUser), nil
}

func (r *userRepo) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	user, err := r.queries.GetUserByID(ctx, int32(id))
	if err != nil {
		return nil, err
	}
	return toDomain(&user), nil
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return toDomain(&user), nil
}

func (r *userRepo) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	users, err := r.queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	return toDomainList(users), nil
}

func (r *userRepo) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	updatedUser, err := r.queries.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:    int32(user.ID),
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		return nil, err
	}
	return toDomain(&updatedUser), nil
}

func (r *userRepo) DeleteUser(ctx context.Context, id int) error {
	return r.queries.DeleteUser(ctx, int32(id))
}

func toDomain(user *sqlc.User) *model.User {
	return &model.User{
		ID:       int(user.ID),
		Email:    user.Email,
		Password: user.Password,
	}
}

func toDomainList(users []sqlc.User) []*model.User {
	var domainUsers []*model.User
	for _, user := range users {
		domainUsers = append(domainUsers, toDomain(&user))
	}
	return domainUsers
}
