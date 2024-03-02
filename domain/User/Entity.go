package user

import (
	"time"
)

type UserRole string

const (
	UserRoleAdmin            UserRole = "admin"
	UserRoleOwner            UserRole = "owner"
	UserRoleSalesManager     UserRole = "sales_manager"
	UserRoleWarehouseManager UserRole = "warehouse_manager"
	UserRoleSalesAgent       UserRole = "sales_agent"
)

func (e *UserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserRole(s)
	case string:
		*e = UserRole(s)
	default:
		return &InvalidUserRole{Message: "Invalid user role"}
	}
	return nil
}

type User struct {
	id        int
	email     string
	password  string
	createdAt time.Time
	firstName string
	lastName  string
	role      UserRole
}

func NewUser(
	email string,
	password string,
	firstName string,
	lastName string,
	role UserRole,
) (*User, error) {

	roleError := role.Scan(role)

	if roleError != nil {
		return &User{}, roleError
	}

	newUser := User{
		id:        -1,
		email:     email,
		password:  password,
		createdAt: time.Now(),
		firstName: firstName,
		lastName:  lastName,
		role:      role,
	}

	return &newUser, nil
}

func (user *User) ID() int {
	return user.id
}

func (user *User) Email() string {
	return user.email
}

func (user *User) Password() string {
	return user.password
}

func (user *User) CreatedAt() time.Time {
	return user.createdAt
}

func (user *User) FirstName() string {
	return user.firstName
}

func (user *User) LastName() string {
	return user.lastName
}

func (user *User) Role() UserRole {
	return user.role
}
