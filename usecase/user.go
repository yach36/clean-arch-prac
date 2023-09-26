package usecase

import (
	"context"

	"github.com/yach36/clean-arch-prac/domain/model"
	"github.com/yach36/clean-arch-prac/domain/repository"
)

type IUserUsecase interface {
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	GetUser(ctx context.Context, id int) (*model.User, error)
	RegisterUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id int) error
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

func (u *userUsecase) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	users, err := u.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userUsecase) GetUser(ctx context.Context, id int) (*model.User, error) {
	user, err := u.repo.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) RegisterUser(ctx context.Context, user *model.User) error {
	err := u.repo.Add(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) DeleteUser(ctx context.Context, id int) error {
	err := u.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
