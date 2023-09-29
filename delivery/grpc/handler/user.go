package handler

import (
	"context"

	_ "github.com/lib/pq"
	userProto "github.com/yach36/clean-arch-prac/delivery/grpc/proto"
	"github.com/yach36/clean-arch-prac/domain/model"
	"github.com/yach36/clean-arch-prac/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewUserServerGrpc(gserver *grpc.Server, userUsecase usecase.IUserUsecase) {
	userServer := &server{
		usecase: userUsecase,
	}

	userProto.RegisterUserServiceServer(gserver, userServer)
	reflection.Register(gserver)
}

type server struct {
	usecase usecase.IUserUsecase
	userProto.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, in *userProto.GetUserRequest) (*userProto.User, error) {
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

func (s *server) transformUserRPC(user *model.User) *userProto.User {
	return &userProto.User{
		ID:   int64(user.ID),
		Name: user.Name,
		Age:  int64(user.Age),
	}
}
