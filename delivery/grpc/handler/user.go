package handler

import (
	"context"

	"github.com/yach36/clean-arch-prac/delivery/grpc/user_grpc"
	"github.com/yach36/clean-arch-prac/domain/model"
	"github.com/yach36/clean-arch-prac/usecase"
	"github.com/yach36/clean-arch-prac/utils/cerrors"
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

var _ user_grpc.UserServiceServer = (*server)(nil)

func (s *server) GetUser(ctx context.Context, in *user_grpc.GetUserRequest) (*user_grpc.User, error) {
	id := 0
	if in != nil {
		id = int(in.Id)
	}

	user, err := s.usecase.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return transformUserRPC(user), err
}

func (s *server) GetUserList(ctx context.Context, in *user_grpc.GetUserListRequest) (*user_grpc.UserList, error) {
	users, err := s.usecase.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	rpcUsers := make([]*user_grpc.User, 0)
	for _, u := range users {
		rpcUsers = append(rpcUsers, transformUserRPC(u))
	}
	result := &user_grpc.UserList{
		Users: rpcUsers,
	}
	return result, nil
}

func (s *server) RegisterUser(ctx context.Context, in *user_grpc.RegisterUserRequest) (*user_grpc.Response, error) {
	user := transformRegisterUserData(in)
	if err := s.usecase.RegisterUser(ctx, user); err != nil {
		return NewResponse(int64(cerrors.StatusCode(err)), "cannot register user"), err
	}
	return NewResponse(200, "success"), nil
}

func transformRegisterUserData(in *user_grpc.RegisterUserRequest) *model.User {
	return &model.User{
		Name: in.GetName(),
		Age:  int(in.GetAge()),
	}
}

func transformUserRPC(user *model.User) *user_grpc.User {
	return &user_grpc.User{
		ID:   int64(user.ID),
		Name: user.Name,
		Age:  int64(user.Age),
	}
}

func NewResponse(status int64, message string) *user_grpc.Response {
	return &user_grpc.Response{
		Status:  status,
		Message: message,
	}
}
