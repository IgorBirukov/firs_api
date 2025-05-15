package userService

import (
	"fmt"
	"net/http"
	"reflect"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(task User) (User, error)
	GetAllUser() ([]User, error)
	UpdateUserByID(id uint, user interface{}) (User, error)
	DeleteUserByID(id uint) (res int, err error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetAllUser() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) UpdateUserByID(id uint, user interface{}) (User, error) {

	fmt.Println(user)
	var fUser User
	err := r.db.First(&fUser, id).Error
	fmt.Println(reflect.TypeOf(user))

	if err != nil {
		return User{}, err
	}

	result := r.db.Model(&fUser).Updates(user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return fUser, nil

}

func (r *userRepository) DeleteUserByID(id uint) (res int, err error) {
	var fuser User
	erro := r.db.First(&fuser, id)
	fmt.Println(&fuser)
	fmt.Println(fuser)
	if erro.Error != nil {
		return http.StatusNotFound, erro.Error
	}
	fmt.Println(1)
	result := r.db.Delete(&fuser)
	fmt.Println(result)
	if result.Error != nil {
		return http.StatusNotFound, result.Error
	}
	return http.StatusNoContent, result.Error
}
