package migration

import (
	"gokil/entity"

	"gorm.io/gorm"
)

type Migration struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Migration {
	return &Migration{
		db: db,
	}
}

func (m Migration) Generate() error {
	err := m.db.Migrator().CreateTable(
		&entity.Users{},
		&entity.Jwt{},
		&entity.Avatars{},
	)

	m.db.Migrator().CreateConstraint(&entity.Users{}, "Jwt")
	m.db.Migrator().CreateConstraint(&entity.Users{}, "Avatars")

	return err
}
