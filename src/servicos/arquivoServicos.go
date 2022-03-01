package servicos

import (
	"gitbuh.com/peevieira/feiras-api/src/modelos/dominios"
	"gitbuh.com/peevieira/feiras-api/src/repositorios"
)

type ArquivoServico struct {
	Repositorio repositorios.ArquivoRepositorio
}

func (serv ArquivoServico) Carregar() *ArquivoServico {
	serv.Repositorio = repositorios.ArquivoRepositorio{}
	return &serv
}

func (serv ArquivoServico) ParaFeira() *[]dominios.Feira {
	feirasArquivo := serv.Repositorio.LerLinhas()
	feiras := []dominios.Feira{}

	for _, feira := range *feirasArquivo {
		feiras = append(feiras,
			dominios.Feira{
				ID:                feira.ID,
				DistritoID:        feira.CodDist,
				DistritoNome:      feira.Distrito,
				SubPrefeituraID:   feira.CodSubPref,
				SubPrefeituraNome: feira.SubPrefe,
				Nome:              feira.Nome_Feira,
				Longitude:         feira.Long,
				Latitude:          feira.Lat,
				Area:              feira.AreaP,
				Censo:             feira.SetCens,
				Regiao5:           feira.Regiao5,
				Regiao8:           feira.Regiao8,
				Registro:          feira.Registro,
				Logradouro:        feira.Logradouro,
				Numero:            feira.Numero,
				Bairro:            feira.Bairro,
				Referencia:        feira.Referencia,
			},
		)
	}

	return &feiras
}
