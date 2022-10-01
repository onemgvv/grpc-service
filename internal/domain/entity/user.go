package entity

import userModel "github.com/onemgvv/grpc-contract/gen/go/users/model/v1"

type User struct {
	Id    int64
	Name  string
	Email string
	Age   uint64
}

func (u *User) ToProto() *userModel.User {
	return &userModel.User{
		Id:    u.Id,
		Name:  u.Name,
		Email: u.Email,
		Age:   u.Age,
	}
}
