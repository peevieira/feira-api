package modelos

import "gorm.io/gorm"

type SubPrefeitura struct {
	gorm.Model
	ID   int64 `gorm:"primary_key,unique"`
	Nome string
}
