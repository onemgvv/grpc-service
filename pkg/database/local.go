package database

import (
	"github.com/onemgvv/grpc-service/internal/delivery/grpc/v1/dto"
	"github.com/onemgvv/grpc-service/internal/domain/entity"
	"sync"
)

type Storage struct {
	sync.RWMutex
	users []*entity.User
}

func NewStorage() *Storage {
	return &Storage{
		users: []*entity.User{},
	}
}

func (s *Storage) SetUser(dto dto.CreateUserDTO) int64 {
	newId := int64(len(s.users))
	s.Lock()
	s.users = append(s.users, &entity.User{Id: newId, Name: *dto.Name, Email: *dto.Email, Age: uint64(*dto.Age)})
	s.Unlock()
	return newId
}

func (s *Storage) UpdateUser(id int64, dto dto.UpdateUserDTO) {
	s.Lock()
	s.users[id] = &entity.User{Id: id, Name: *dto.Name, Email: *dto.Email, Age: uint64(*dto.Age)}
	s.Unlock()
}

func (s *Storage) GetUser(id int64) *entity.User {
	s.RLock()
	defer s.RUnlock()
	return s.users[id]
}

func (s *Storage) GetUsers() []*entity.User {
	s.RLock()
	defer s.RUnlock()
	return s.users
}
