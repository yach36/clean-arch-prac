package repository

import (
	"context"

	"github.com/yach36/clean-arch-prac/domain/model"
)

type IUserRepository interface {
	FindAll(ctx context.Context) ([]*model.User, error)
	Find(ctx context.Context, id int) (*model.User, error)
}
