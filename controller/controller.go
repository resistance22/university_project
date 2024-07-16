package controller

import usecase "github.com/resistance22/university_project/usecase"

func NewUserController(usecase *usecase.UserUseCase) *userController {
	return &userController{
		UseCases: usecase,
	}
}
