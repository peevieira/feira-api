package repositorios

import (
	"gitbuh.com/peevieira/feiras-api/src/modelos"
	"gitbuh.com/peevieira/feiras-api/src/modelos/dominios"
	errorStack "github.com/pkg/errors"
	"gorm.io/gorm"
)

type FeiraRepositorio struct {
	DB *gorm.DB
}

func (repo FeiraRepositorio) Criar(feira *modelos.Feira) (*dominios.Feira, error) {
	DB := repo.DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(feira)

	if DB.Error != nil {
		return nil, errorStack.WithStack(DB.Error)
	}

	return dominios.ParaFeiraDominio(feira), nil
}

func (repo FeiraRepositorio) Recuperar(ID int64) (*dominios.Feira, error) {
	var f modelos.Feira
	DB := repo.DB.Preload("Distrito").Preload("SubPrefeitura").First(&f, ID)

	if DB.Error != nil {
		return nil, errorStack.WithStack(DB.Error)
	}

	return dominios.ParaFeiraDominio(&f), nil
}

func (repo FeiraRepositorio) RecuperarComFiltro(campo string, valor string) (*dominios.Feira, error) {
	var f modelos.Feira
	DB := repo.DB.Preload("Distrito").Preload("SubPrefeitura").Joins("LEFT JOIN DISTRITOS ON FEIRAS.DISTRITO_ID = DISTRITOS.ID").Where(repo.MontarQuery(campo), repo.MontarCondicao(valor)).First(&f)

	if DB.Error != nil {
		return nil, errorStack.WithStack(DB.Error)
	}

	return dominios.ParaFeiraDominio(&f), nil
}

func (repo FeiraRepositorio) Atualizar(f *modelos.Feira) (*dominios.Feira, error) {
	DB := repo.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(f)

	if DB.Error != nil {
		return nil, errorStack.WithStack(DB.Error)
	}

	return dominios.ParaFeiraDominio(f), nil
}

func (repo FeiraRepositorio) Deletar(feira *modelos.Feira) (*dominios.Feira, error) {
	DB := repo.DB.Delete(feira, feira.ID)

	if DB.Error != nil {
		return nil, errorStack.WithStack(DB.Error)
	}

	return dominios.ParaFeiraDominio(feira), nil
}

func (repo FeiraRepositorio) RecuperarTudo() (*[]dominios.Feira, error) {
	var f []modelos.Feira
	DB := repo.DB.Preload("Distrito").Preload("SubPrefeitura").Find(&f)

	if DB.Error != nil {
		return nil, errorStack.WithStack(DB.Error)
	}

	return dominios.ParaListaFeiraDominio(&f), nil
}

func (repo FeiraRepositorio) MontarQuery(campo string) string {
	if campo == "regiao5" {
		return "regiao5 LIKE ? "
	} else if campo == "nome" {
		return "feiras.nome LIKE ? "
	} else if campo == "bairro" {
		return "bairro LIKE ? "
	} else if campo == "distrito" {
		return "distritos.nome LIKE ? "
	}

	return ""
}

func (repo FeiraRepositorio) MontarCondicao(valor string) string {
	return "%" + valor + "%"
}
