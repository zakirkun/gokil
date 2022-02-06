package repository

import (
	"gokil/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (entity.Users, error)
	FindByID(id int) (entity.Users, error)
	Save(user entity.Users) (entity.Users, error)
	Update(user entity.Users) (entity.Users, error)
}

type userRepository struct {
	db gorm.DB
}

func NewUserRepository(db gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r userRepository) FindByEmail(email string) (entity.Users, error) {
	var user entity.Users
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r userRepository) FindByID(id int) (entity.Users, error) {
	var user entity.Users
	err := r.db.Where("id = ?", id).First(&user).Error
	return user, err
}

func (r userRepository) Save(user entity.Users) (entity.Users, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r userRepository) Update(user entity.Users) (entity.Users, error) {
	err := r.db.Save(&user).Error
	return user, err
}
