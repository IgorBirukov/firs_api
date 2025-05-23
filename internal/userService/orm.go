package userService

import (
	userService "first_api/internal/taskService"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Tasks    []userService.Task `json:"tasks" gorm:"foreignKey:UserID"`
}
