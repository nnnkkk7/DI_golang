package service

import (
	"context"

	"github.com/nnnkkk7/DI_golang/model"
	"github.com/nnnkkk7/DI_golang/repository"
)

type UserService interface {
	CreateUser(c context.Context, user *model.User)
}
type userService struct {
	ur repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{ur}
}

func (u *userService) CreateUser(c context.Context, user *model.User) {
	u.ur.CreateUser(c, user)
}
