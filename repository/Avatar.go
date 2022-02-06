package repository

import (
	"gokil/entity"

	"gorm.io/gorm"
)

type AvatarRepository interface {
	Save(avatar entity.Avatars) (entity.Avatars, error)
	Get(id int) (entity.Avatars, error)
	Update(avatar entity.Avatars) (entity.Avatars, error)
}

type avatarRepository struct {
	db gorm.DB
}

func NewAvatarRepository(db gorm.DB) *avatarRepository {
	return &avatarRepository{db: db}
}

func (r avatarRepository) Save(avatar entity.Avatars) (entity.Avatars, error) {
	err := r.db.Create(&avatar).Error
	return avatar, err
}

func (r avatarRepository) Get(id int) (entity.Avatars, error) {
	var avatar entity.Avatars
	err := r.db.Where("id = ?", id).First(&avatar).Error
	return avatar, err
}

func (r avatarRepository) Update(avatar entity.Avatars) (entity.Avatars, error) {
	err := r.db.Save(&avatar).Error
	return avatar, err
}
