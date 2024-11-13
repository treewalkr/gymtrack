// internal/infrastructure/persistence/user_repository.go
package persistence

import (
	"github.com/treewalkr/gymtrack/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(dsn string) (domain.UserRepository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		return nil, err
	}

	return &UserRepository{db: db}, nil
}

// Ensure UserRepository implements domain.UserRepository
var _ domain.UserRepository = &UserRepository{}

func (r *UserRepository) CreateUser(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetUserByID(id string) (*domain.User, error) {
	var user domain.User
	result := r.db.First(&user, "id = ?", id)
	return &user, result.Error
}

func (r *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	result := r.db.First(&user, "email = ?", email)
	return &user, result.Error
}

func (r *UserRepository) UpdateUser(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) DeleteUser(id string) error {
	return r.db.Delete(&domain.User{}, "id = ?", id).Error
}
