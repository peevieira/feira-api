package modelos

import (
	"gorm.io/gorm"
)

type Feira struct {
	gorm.Model
	ID              int64 `gorm:"primary_key"`
	DistritoID      int64
	Distrito        Distrito `gorm:"foreignKey:DistritoID;references:ID"`
	SubPrefeituraID int64
	SubPrefeitura   SubPrefeitura `gorm:"foreignKey:SubPrefeituraID;references:ID"`
	Longitude       float64
	Latitude        float64
	Censo           int64
	Area            float64
	Regiao5         string
	Regiao8         string
	Nome            string
	Registro        string
	Logradouro      string
	Numero          string
	Bairro          string
	Referencia      string
}
