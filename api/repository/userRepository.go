package repository

import (
	"errors"
	"fmt"

	"dart-api/api/model"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Create(u *model.User) error
	Update(u *model.User) error
	GetAll() ([]*model.User, error)
	GetByID(id string) (*model.User, error)
	Delete(id string) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

// get all users
func (u *userRepo) GetAll() ([]*model.User, error) {
	// put logs here
	users := make([]*model.User, 0)
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Create a user
func (u *userRepo) Create(usr *model.User) error {
	// put logs here
	err := u.db.Create(&usr).Error
	if err != nil {
		return err
	}
	return nil
}

// find a User
func (u *userRepo) GetByID(id string) (*model.User, error) {
	user := &model.User{}
	err := u.db.Where("user_id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

//  Update user details
func (u *userRepo) Update(usr *model.User) error {
	err := u.db.Model(&usr).Update(model.User{
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		UserName:  usr.UserName,
		Password:  usr.Password,
		Email:     usr.Email}).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete User Details
func (u *userRepo) Delete(id string) error {
	if u.db.Delete(&model.User{}, "user_id = ?", id).Error != nil {
		errMsg := fmt.Sprintf("error while deleting the user with id : %s", id)
		return errors.New(errMsg)
	}
	return nil
}
