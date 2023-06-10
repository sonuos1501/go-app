package services

import (
	"github.com/sonuos1501/go-app/adapter/dtos"
	"github.com/sonuos1501/go-app/public/resources"
)

type IUserService interface {
	CreateUser(data dtos.SignupDto) (*resources.UserResource, error)
	Login(data dtos.LoginDto) (*resources.LoginResource, error)
	FindById(id uint) (*resources.UserResource, error)
}
