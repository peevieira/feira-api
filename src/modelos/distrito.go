package modelos

import "gorm.io/gorm"

type Distrito struct {
	gorm.Model
	ID   int64 `gorm:"primary_key"`
	Nome string
}
