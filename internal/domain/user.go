// internal/domain/user.go
package domain

type User struct {
	// uuid
	ID        string
	Username  string
	Email     string
	Password  string
	Role      string // e.g., "trainer", "client"
	CreatedAt string
	UpdatedAt string
}

type UserRepository interface {
	CreateUser(user *User) error
	GetUserByID(id string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
}
