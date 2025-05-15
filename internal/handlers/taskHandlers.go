package handlers

import (
	"context"
	"first_api/internal/taskService"
	"first_api/internal/web/tasks"
	"fmt"
)

type Handler struct {
	Service *taskService.TaskService
}

// GetTasksId implements tasks.StrictServerInterface.
func (h *Handler) GetTasksId(ctx context.Context, request tasks.GetTasksIdRequestObject) (tasks.GetTasksIdResponseObject, error) {
	allTasks, err := h.Service.GetTasksByUserID(uint(request.Id))
	if err != nil {
		return nil, err
	}

	responce := tasks.GetTasksId200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		responce = append(responce, task)
	}

	return responce, nil
}

// DeleteTasksId implements tasks.StrictServerInterface.
func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {

	resDell, err := h.Service.DeleteTaskByID(uint(request.Id))

	if err != nil {
		return nil, err
	}

	if resDell == 204 {
		response := tasks.DeleteTasksId204Response{}
		return response, nil
	}
	return tasks.DeleteTasksId400Response{}, nil
}

// PatchTasksId implements tasks.StrictServerInterface.
func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {

	if request.Id == 0 {
		return nil, fmt.Errorf("id, task, or is_done cannot be nil")
	}

	taskRequest := request.Body
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	UpdateTask, err := h.Service.UpdateTaskByID(uint(request.Id), taskToCreate)

	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &UpdateTask.ID,
		Task:   &UpdateTask.Task,
		IsDone: &UpdateTask.IsDone,
	}

	return response, nil
}

/* Старый гет
// GetTasks implements tasks.StrictServerInterface.
func (h *Handler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTask()
	if err != nil {
		return nil, err
	}

	responce := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		responce = append(responce, task)
	}

	return responce, nil
}
*/

// GetTasks implements tasks.StrictServerInterface.
func (h *Handler) GetTasksByUserID(ctx context.Context, request tasks.GetTasksIdRequestObject) (tasks.GetTasksIdResponseObject, error) {
	allTasks, err := h.Service.GetTasksByUserID(uint(request.Id))
	if err != nil {
		return nil, err
	}

	responce := tasks.GetTasksId200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		responce = append(responce, task)
	}

	return responce, nil
}

// PostTasks implements tasks.StrictServerInterface.
func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,

		//+ 18.04.2025
		UserID: *taskRequest.UserId,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
		//+ 18.04.2025
		UserId: &createdTask.UserID,
	}

	return response, nil
}

//func (h *Handler) GetTasksByUserID(ctx context.Context, request task.Get)

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}
