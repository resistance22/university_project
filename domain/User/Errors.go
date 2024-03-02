package user

type InvalidUserRole struct {
	Message string
}

func NewInvalidUserRoleError(msg string) *InvalidUserRole {
	return &InvalidUserRole{msg}
}

func (e *InvalidUserRole) Error() string {
	return e.Message
}
