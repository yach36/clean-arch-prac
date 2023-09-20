package usecase

import (
	"context"

	"github.com/yach36/clean-arch-prac/domain/model"
	"github.com/yach36/clean-arch-prac/domain/repository"
)

type IUserUsecase interface {
	GetAll(ctx context.Context) ([]*model.User, error)
	Get(ctx context.Context, id int) (*model.User, error)
}

type userUsecase struct {
	repo repository.IUserRepository
}

var _ IUserUsecase = (*userUsecase)(nil)

func NewUserUsecase(r repository.IUserRepository) IUserUsecase {
	return &userUsecase{
		repo: r,
	}
}

func (u *userUsecase) GetAll(ctx context.Context) ([]*model.User, error) {
	users, err := u.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userUsecase) Get(ctx context.Context, id int) (*model.User, error) {
	user, err := u.repo.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
