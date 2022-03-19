package service

import (
	helper "github.com/nhatlang19/go-monorepo/pkg/helper"
	grpc_client "github.com/nhatlang19/go-monorepo/services/user/client"
	"github.com/nhatlang19/go-monorepo/services/user/model"
	"github.com/nhatlang19/go-monorepo/services/user/repository"

	"log"
)

type UserService interface {
	Save(model.User) (model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
	mailClient     grpc_client.MailClient
}

func NewUserService(r repository.UserRepository, m grpc_client.MailClient) UserService {
	return userService{
		userRepository: r,
		mailClient:     m,
	}
}

func (u userService) Save(user model.User) (model.User, error) {
	log.Print("[UserService]...Save")
	user.Password, _ = helper.HashPassword(user.Password)
	user, err := u.userRepository.Save(user)

	u.mailClient.HandleRegisterMail(user)

	return user, err
}
