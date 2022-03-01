package dominios

import "gitbuh.com/peevieira/feiras-api/src/modelos"

// Feira represents the model for an feira
type Feira struct {
	ID                int64
	DistritoID        int64
	DistritoNome      string `validate:"required"`
	SubPrefeituraID   int64
	SubPrefeituraNome string  `validate:"required"`
	Longitude         float64 `validate:"required"`
	Latitude          float64 `validate:"required"`
	Censo             int64   `validate:"required"`
	Area              float64 `validate:"required"`
	Regiao5           string  `validate:"required"`
	Regiao8           string  `validate:"required"`
	Nome              string  `validate:"required"`
	Registro          string  `validate:"required"`
	Logradouro        string  `validate:"required"`
	Numero            string  `validate:"required"`
	Bairro            string  `validate:"required"`
	Referencia        string  `validate:"required"`
}

func (df *Feira) ParaModelo() *modelos.Feira {
	return &modelos.Feira{
		ID:         df.ID,
		DistritoID: df.DistritoID,
		Distrito: modelos.Distrito{
			ID:   df.DistritoID,
			Nome: df.DistritoNome,
		},
		SubPrefeituraID: df.DistritoID,
		SubPrefeitura: modelos.SubPrefeitura{
			ID:   df.SubPrefeituraID,
			Nome: df.SubPrefeituraNome,
		},
		Longitude:  df.Longitude,
		Latitude:   df.Latitude,
		Censo:      df.Censo,
		Area:       df.Area,
		Regiao5:    df.Regiao5,
		Regiao8:    df.Regiao8,
		Nome:       df.Nome,
		Registro:   df.Registro,
		Logradouro: df.Logradouro,
		Numero:     df.Numero,
		Bairro:     df.Bairro,
		Referencia: df.Referencia,
	}
}

func ParaFeiraDominio(mf *modelos.Feira) *Feira {
	return &Feira{
		ID:                mf.ID,
		DistritoID:        mf.Distrito.ID,
		DistritoNome:      mf.Distrito.Nome,
		SubPrefeituraID:   mf.SubPrefeitura.ID,
		SubPrefeituraNome: mf.SubPrefeitura.Nome,
		Longitude:         mf.Longitude,
		Latitude:          mf.Latitude,
		Censo:             mf.Censo,
		Area:              mf.Area,
		Regiao5:           mf.Regiao5,
		Regiao8:           mf.Regiao8,
		Nome:              mf.Nome,
		Registro:          mf.Registro,
		Logradouro:        mf.Logradouro,
		Numero:            mf.Numero,
		Bairro:            mf.Bairro,
		Referencia:        mf.Referencia,
	}
}

func ParaListaFeiraDominio(l *[]modelos.Feira) *[]Feira {
	feirasModelo := *l
	var feiras []Feira

	for _, feira := range feirasModelo {
		feiras = append(feiras, *ParaFeiraDominio(&feira))
	}

	return &feiras
}
