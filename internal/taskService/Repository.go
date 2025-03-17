package taskService

import (
	"fmt"
	"net/http"
	"reflect"

	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTask() ([]Task, error)
	UpdateTaskByID(id uint, task interface{}) (Task, error)
	DeleteTaskByID(id uint) (res int, err error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTask() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, task interface{}) (Task, error) {

	fmt.Println(task)
	var fTask Task
	err := r.db.First(&fTask, id).Error
	fmt.Println(reflect.TypeOf(task))
	//fmt.Println(reflect.TypeOf(taskService.Task))
	if err != nil {
		return Task{}, err
	}

	fmt.Println(&fTask)
	result := r.db.Model(&fTask).Updates(task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return fTask, nil

}

func (r *taskRepository) DeleteTaskByID(id uint) (res int, err error) {
	var ftask Task
	erro := r.db.First(&ftask, id)
	fmt.Println(&ftask)
	fmt.Println(ftask)
	if erro.Error != nil {
		return http.StatusNotFound, erro.Error
	}
	fmt.Println(1)
	result := r.db.Delete(&ftask)
	fmt.Println(result)
	if result.Error != nil {
		return http.StatusNotFound, result.Error
	}
	return http.StatusNoContent, result.Error
}
