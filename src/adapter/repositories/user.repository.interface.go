package repositories

import (
	"github.com/sonuos1501/go-app/adapter/dtos"
	"github.com/sonuos1501/go-app/adapter/models"
)

type IUserRepository interface {
	CreateUser(data dtos.SignupDto) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindById(id uint) (*models.User, error)
}
