package repository

import (
	"gokil/entity"

	"gorm.io/gorm"
)

type JwtRepository interface {
	Save(jwt entity.Jwt) (entity.Jwt, error)
	GetRefreshToken(refreshToken string) (entity.Jwt, error)
	Delete(refreshToken string) error
}

type jwtRepository struct {
	db gorm.DB
}

func NewJwtRepository(db gorm.DB) *jwtRepository {
	return &jwtRepository{db: db}
}

func (r jwtRepository) Save(jwt entity.Jwt) (entity.Jwt, error) {
	err := r.db.Create(&jwt).Error
	return jwt, err
}

func (r jwtRepository) GetRefreshToken(refreshToken string) (entity.Jwt, error) {
	var jwt entity.Jwt
	err := r.db.Where("refresh_token = ?", refreshToken).First(&jwt).Error
	return jwt, err
}

func (r jwtRepository) Delete(refreshToken string) error {
	err := r.db.Where("refresh_token = ?", refreshToken).Delete(&entity.Jwt{}).Error
	return err
}
