package mocks

import (
	"gitbuh.com/peevieira/feiras-api/src/modelos"
	"gitbuh.com/peevieira/feiras-api/src/modelos/dominios"
	"gitbuh.com/peevieira/feiras-api/src/servicos"
	"gorm.io/gorm"
)

type FeiraMock struct {
	DB *gorm.DB
}

func (mock FeiraMock) CriarUmaFeira() *modelos.Feira {

	f := modelos.Feira{
		ID:         mock.IdRegistroTest(),
		Longitude:  9999.9,
		Latitude:   9999.9,
		Censo:      9999,
		Area:       9999.9,
		Regiao5:    "Região 5 mock",
		Regiao8:    "Região 8 mock",
		Nome:       "Feira mock",
		Registro:   "Registro mock",
		Logradouro: "Logradouro mock",
		Numero:     "S/N mock",
		Bairro:     "Bairro mock",
		Referencia: "Referencia mock",
	}

	mock.DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(f)

	return &f
}

func (mock FeiraMock) CriarUmaFeiraParaFiltrarRegiao5() *dominios.Feira {

	service := servicos.FeiraServico{}.Carregar()

	f := dominios.Feira{
		ID:                -10,
		DistritoID:        -10,
		DistritoNome:      "Distrino",
		SubPrefeituraID:   -10,
		SubPrefeituraNome: "Sub prefeitura",
		Longitude:         0.0,
		Latitude:          0.0,
		Censo:             0,
		Area:              0.0,
		Regiao5:           "filtro",
		Regiao8:           "Regiao 8",
		Nome:              "Nome",
		Registro:          "Registro",
		Logradouro:        "Logradouro",
		Numero:            "S/N",
		Bairro:            "Bairro",
		Referencia:        "Referencia",
	}

	r, _ := service.Criar(&f)

	return r
}

func (mock FeiraMock) CriarUmaFeiraParaFiltrarNome() *dominios.Feira {

	service := servicos.FeiraServico{}.Carregar()

	f := dominios.Feira{
		ID:                -20,
		DistritoID:        -20,
		DistritoNome:      "Distrito",
		SubPrefeituraID:   -20,
		SubPrefeituraNome: "Sub prefeitura",
		Longitude:         0.0,
		Latitude:          0.0,
		Censo:             0,
		Area:              0.0,
		Regiao5:           "Regiao 5",
		Regiao8:           "Regiao 8",
		Nome:              "filtro",
		Registro:          "Registro",
		Logradouro:        "Logradouro",
		Numero:            "S/N",
		Bairro:            "Bairro",
		Referencia:        "Referencia",
	}

	r, _ := service.Criar(&f)

	return r
}

func (mock FeiraMock) CriarUmaFeiraParaFiltrarBairro() *dominios.Feira {

	service := servicos.FeiraServico{}.Carregar()

	f := dominios.Feira{
		ID:                -30,
		DistritoID:        -30,
		DistritoNome:      "Distrito",
		SubPrefeituraID:   -30,
		SubPrefeituraNome: "Sub prefeitura",
		Longitude:         0.0,
		Latitude:          0.0,
		Censo:             0,
		Area:              0.0,
		Regiao5:           "Regiao 5",
		Regiao8:           "Regiao 8",
		Nome:              "Nome",
		Registro:          "Registro",
		Logradouro:        "Logradouro",
		Numero:            "S/N",
		Bairro:            "filtro",
		Referencia:        "Referencia",
	}

	r, _ := service.Criar(&f)

	return r
}

func (mock FeiraMock) CriarUmaFeiraParaFiltrarDistrito() *dominios.Feira {

	service := servicos.FeiraServico{}.Carregar()

	f := dominios.Feira{
		ID:                -40,
		DistritoID:        -40,
		DistritoNome:      "filtro",
		SubPrefeituraID:   -40,
		SubPrefeituraNome: "Sub prefeitura",
		Longitude:         0.0,
		Latitude:          0.0,
		Censo:             0,
		Area:              0.0,
		Regiao5:           "Regiao 5",
		Regiao8:           "Regiao 8",
		Nome:              "Nome",
		Registro:          "Registro",
		Logradouro:        "Logradouro",
		Numero:            "S/N",
		Bairro:            "Bairro",
		Referencia:        "Referencia",
	}

	r, _ := service.Criar(&f)

	return r
}

func (mock FeiraMock) DeletarUmaFeira() *modelos.Feira {
	f := modelos.Feira{}

	mock.DB.Session(&gorm.Session{FullSaveAssociations: true}).Delete(&f, mock.IdRegistroTest)

	return &f
}

func (mock FeiraMock) CriarDominioFeira() *dominios.Feira {
	return &dominios.Feira{
		ID:                mock.IdRegistroTest(),
		DistritoNome:      "Distrito mock",
		SubPrefeituraNome: "Sub prefeitura mock",
		Longitude:         888.8,
		Latitude:          888.8,
		Censo:             888,
		Area:              888.8,
		Regiao5:           "Região 5",
		Regiao8:           "Região 8",
		Nome:              "Feira mock de teste",
		Registro:          "Registro",
		Logradouro:        "Logradouro",
		Numero:            "Numero",
		Bairro:            "Bairro",
		Referencia:        "Referencia",
	}
}

func (mock FeiraMock) CriarDominioFeiraSemValor(c string) *dominios.Feira {
	f := mock.CriarDominioFeira()

	if c == "dist" {
		f.DistritoNome = ""
	} else if c == "pref" {
		f.SubPrefeituraNome = ""
	} else if c == "long" {
		f.Longitude = 0
	} else if c == "lat" {
		f.Latitude = 0
	} else if c == "censo" {
		f.Censo = 0
	} else if c == "area" {
		f.Area = 0
	} else if c == "reg5" {
		f.Regiao5 = ""
	} else if c == "reg8" {
		f.Regiao8 = ""
	} else if c == "nome" {
		f.Nome = ""
	} else if c == "registro" {
		f.Registro = ""
	} else if c == "logra" {
		f.Logradouro = ""
	} else if c == "num" {
		f.Numero = ""
	} else if c == "bairro" {
		f.Bairro = ""
	} else {
		f.Referencia = ""
	}

	return f
}

func (mock FeiraMock) RespostaValidador(c string) string {
	if c == "dist" {
		return "[{\"Namespace\":\"Feira.DistritoNome\",\"Field\":\"DistritoNome\",\"Tag\":\"required\",\"Type\":\"string\"}]"
	} else if c == "pref" {
		return "[{\"Namespace\":\"Feira.SubPrefeituraNome\",\"Field\":\"SubPrefeituraNome\",\"Tag\":\"required\",\"Type\":\"string\"}]"
	} else if c == "long" {
		return "[{\"Namespace\":\"Feira.Longitude\",\"Field\":\"Longitude\",\"Tag\":\"required\",\"Type\":\"float64\"}]"
	} else if c == "lat" {
		return "[{\"Namespace\":\"Feira.Latitude\",\"Field\":\"Latitude\",\"Tag\":\"required\",\"Type\":\"float64\"}]"
	} else if c == "censo" {
		return "[{\"Namespace\":\"Feira.Censo\",\"Field\":\"Censo\",\"Tag\":\"required\",\"Type\":\"int64\"}]"
	} else if c == "area" {
		return "[{\"Namespace\":\"Feira.Area\",\"Field\":\"Area\",\"Tag\":\"required\",\"Type\":\"float64\"}]"
	} else if c == "reg5" {
		return "[{\"Namespace\":\"Feira.Regiao5\",\"Field\":\"Regiao5\",\"Tag\":\"required\",\"Type\":\"string\"}]"
	} else if c == "reg8" {
		return "[{\"Namespace\":\"Feira.Regiao8\",\"Field\":\"Regiao8\",\"Tag\":\"required\",\"Type\":\"string\"}]"
	} else if c == "nome" {
		return "[{\"Namespace\":\"Feira.Nome\",\"Field\":\"Nome\",\"Tag\":\"required\",\"Type\":\"string\"}]"
	} else if c == "registro" {
		return "[{\"Namespace\":\"Feira.Registro\",\"Field\":\"Registro\",\"Tag\":\"required\",\"Type\":\"string\"}]"
	} else if c == "logra" {
		return "[{\"Namespace\":\"Feira.Logradouro\",\"Field\":\"Logradouro\",\"Tag\":\"required\",\"Type\":\"string\"}]"
	} else if c == "num" {
		return "[{\"Namespace\":\"Feira.Numero\",\"Field\":\"Numero\",\"Tag\":\"required\",\"Type\":\"string\"}]"
	} else if c == "bairro" {
		return "[{\"Namespace\":\"Feira.Bairro\",\"Field\":\"Bairro\",\"Tag\":\"required\",\"Type\":\"string\"}]"
	} else {
		return "[{\"Namespace\":\"Feira.Referencia\",\"Field\":\"Referencia\",\"Tag\":\"required\",\"Type\":\"string\"}]"
	}
}

func (mock FeiraMock) RecuperarUmaFeiraDoBanco() *dominios.Feira {
	var feira modelos.Feira
	mock.DB.Preload("Distrito").Preload("SubPrefeitura").First(&feira, mock.IdRegistroTest())

	return dominios.ParaFeiraDominio(&feira)
}

func (mock FeiraMock) RespostasAoDeletar() string {
	return `{"sucesso":"Registro deletado"}`
}

func (mock FeiraMock) IdRegistroTest() int64 {
	return 999999
}
