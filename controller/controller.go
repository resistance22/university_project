package controller

func NewUserController(usecase UserUseCase) *userController {
	return &userController{
		UseCases: usecase,
	}
}
