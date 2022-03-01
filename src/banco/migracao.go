package banco

import (
	"gitbuh.com/peevieira/feiras-api/src/modelos"
	"gorm.io/gorm"
)

func Migrador(DB *gorm.DB) error {
	err := DB.AutoMigrate(
		&modelos.Distrito{},
		&modelos.SubPrefeitura{},
		&modelos.Feira{},
	)

	if err != nil {
		return err
	}

	return nil
}
