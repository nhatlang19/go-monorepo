package repository

import (
	"log"

	"github.com/nhatlang19/go-monorepo/pkg/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(model.User) (model.User, error)
	GetAll() ([]model.User, error)
	GetByEmail(string) (model.User, error)
	Migrate() error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		DB: db,
	}
}

func (u userRepository) Migrate() error {
	log.Println("[UserRepository]...Migrate")
	return u.DB.AutoMigrate(&model.User{})
}

func (u userRepository) Save(user model.User) (model.User, error) {
	log.Print("[UserRepository]...Save")
	err := u.DB.Create(&user).Error
	return user, err
}

func (u userRepository) GetByEmail(email string) (model.User, error) {
	log.Print("[UserRepository]...GetByEmail")
	var user model.User
	err := u.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func (u userRepository) GetAll() ([]model.User, error) {
	log.Print("[UserRepository]...Get All")
	var users []model.User
	err := u.DB.Find(&users).Error
	return users, err
}
