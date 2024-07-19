package controller

func NewUserController(usecase IUserUseCase) *userController {
	return &userController{
		UseCases: usecase,
	}
}
