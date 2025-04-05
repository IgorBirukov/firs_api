package handlers

import (
	"context"
	"first_api/internal/userService"
	"first_api/internal/web/users"
	"fmt"
)

type HandlerUser struct {
	Service *userService.UserService
}

func (h *HandlerUser) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {

	resDell, err := h.Service.DeleteUserByID(uint(request.Id))

	if err != nil {
		return nil, err
	}

	if resDell == 204 {
		response := users.DeleteUsersId200Response{}
		return response, nil
	}
	return users.DeleteUsersId400Response{}, nil
}

// PatchTasksId implements tasks.StrictServerInterface.
func (h *HandlerUser) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {

	if request.Id == 0 {
		return nil, fmt.Errorf("id, user, or is_done cannot be nil")
	}

	userRequest := request.Body
	userToCreate := userService.User{
		Password: *userRequest.Password,
		Email:    *userRequest.Email,
	}

	UpdateTask, err := h.Service.UpdateUserByID(uint(request.Id), userToCreate)

	if err != nil {
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		Id:       &UpdateTask.ID,
		Password: &UpdateTask.Password,
		Email:    &UpdateTask.Email,
	}

	return response, nil
}

// GetTasks implements tasks.StrictServerInterface.
func (h *HandlerUser) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetAllUser()
	if err != nil {
		return nil, err
	}

	responce := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Password: &usr.Password,
			Email:    &usr.Email,
		}
		responce = append(responce, user)
	}

	return responce, nil
}

// PostTasks implements tasks.StrictServerInterface.
func (h *HandlerUser) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	userToCreate := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	createdUser, err := h.Service.CreateUser(userToCreate)

	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}

	return response, nil
}

func NewHandlerUser(service *userService.UserService) *HandlerUser {
	return &HandlerUser{
		Service: service,
	}
}
