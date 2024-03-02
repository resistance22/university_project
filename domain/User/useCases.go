package user

import db "github.com/resistance22/micorsales/db/sqlc"

type IUseCases interface {
	Register(data db.CreateUserParams) (*User, error)
}

type UseCases struct {
	UserRepository IUserRepository
}

func NewUserUseCase(repository IUserRepository) *UseCases {
	return &UseCases{repository}
}

func (useCases UseCases) Register(data db.CreateUserParams) (*User, error) {
	newUser, err := NewUser(
		data.Email,
		data.Password,
		data.FirstName,
		data.LastName,
		data.Role,
	)

	useCases.UserRepository.createUser(newUser)

}
