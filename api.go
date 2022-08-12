package ggorm

import "gorm.io/gorm"

func NewGormRepository[Belongings any](db *gorm.DB) *GormRepository[Belongings] {
	return &GormRepository[Belongings]{
		db: db,
	}
}
