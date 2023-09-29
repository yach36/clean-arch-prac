package handler

import (
	"context"

	_ "github.com/lib/pq"
	"github.com/yach36/clean-arch-prac/delivery/grpc/user_grpc"
	"github.com/yach36/clean-arch-prac/domain/model"
	"github.com/yach36/clean-arch-prac/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewUserServerGrpc(gserver *grpc.Server, userUsecase usecase.IUserUsecase) {
	userServer := &server{
		usecase: userUsecase,
	}

	user_grpc.RegisterUserServiceServer(gserver, userServer)
	reflection.Register(gserver)
}

type server struct {
	usecase usecase.IUserUsecase
	user_grpc.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, in *user_grpc.GetUserRequest) (*user_grpc.User, error) {
	id := 0
	if in != nil {
		id = int(in.Id)
	}

	user, err := s.usecase.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.transformUserRPC(user), err
}

func (s *server) transformUserRPC(user *model.User) *user_grpc.User {
	return &user_grpc.User{
		ID:   int64(user.ID),
		Name: user.Name,
		Age:  int64(user.Age),
	}
}
