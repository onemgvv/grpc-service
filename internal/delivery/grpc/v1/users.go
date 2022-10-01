package v1

import (
	"context"
	userModel "github.com/onemgvv/grpc-contract/gen/go/users/model/v1"
	userService "github.com/onemgvv/grpc-contract/gen/go/users/service/v1"
	"github.com/onemgvv/grpc-service/internal/delivery/grpc/v1/dto"
	"github.com/onemgvv/grpc-service/pkg/database"
)

type UserServer struct {
	userService.UnimplementedUserServiceServer
	*database.Storage
}

func NewUserServer(
	unimplementedUserServiceServer userService.UnimplementedUserServiceServer,
	storage *database.Storage,
) *UserServer {
	return &UserServer{UnimplementedUserServiceServer: unimplementedUserServiceServer, Storage: storage}
}

func (s *UserServer) CreateUser(
	_ context.Context,
	request *userService.CreateUserRequest,
) (*userService.CreateUserResponse, error) {
	createDto := dto.CreateUserDTO{Name: &request.Name, Email: &request.Email, Age: &request.Age}
	id := s.Storage.SetUser(createDto)
	return &userService.CreateUserResponse{
		User: s.GetUser(id).ToProto(),
	}, nil
}

func (s *UserServer) GetUsers(
	context.Context,
	*userService.GetUsersRequest,
) (*userService.GetUsersResponse, error) {
	var users []*userModel.User

	for _, item := range s.Storage.GetUsers() {
		users = append(users, item.ToProto())
	}

	return &userService.GetUsersResponse{
		Users: users,
	}, nil
}

func (s *UserServer) UpdateUser(
	_ context.Context,
	request *userService.UpdateUserRequest,
) (*userService.UpdateUserResponse, error) {
	updateDto := dto.UpdateUserDTO{Name: &request.Name, Email: &request.Email, Age: &request.Age}
	s.Storage.UpdateUser(request.Id, updateDto)
	return &userService.UpdateUserResponse{
		User: s.GetUser(request.Id).ToProto(),
	}, nil
}
