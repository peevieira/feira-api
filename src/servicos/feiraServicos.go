package servicos

import (
	"gitbuh.com/peevieira/feiras-api/src/banco"
	"gitbuh.com/peevieira/feiras-api/src/modelos/dominios"
	"gitbuh.com/peevieira/feiras-api/src/repositorios"
)

type FeiraServico struct {
	Repositorio repositorios.FeiraRepositorio
}

func (serv FeiraServico) Carregar() *FeiraServico {
	serv.Repositorio = repositorios.FeiraRepositorio{DB: banco.DB}
	return &serv
}

func (serv FeiraServico) Criar(f *dominios.Feira) (*dominios.Feira, error) {
	return serv.Repositorio.Criar(f.ParaModelo())
}

func (serv FeiraServico) CriarObjetoPronto(f *dominios.Feira) (*dominios.Feira, error) {
	return serv.Repositorio.Criar(f.ParaModelo())
}

func (serv FeiraServico) RecuperarComFiltro(campo string, valor string) (*dominios.Feira, error) {
	return serv.Repositorio.RecuperarComFiltro(campo, valor)
}

func (serv FeiraServico) Recuperar(ID int64) (*dominios.Feira, error) {
	return serv.Repositorio.Recuperar(ID)
}

func (serv FeiraServico) Atualizar(f *dominios.Feira, ID int64) (*dominios.Feira, error) {
	return serv.Repositorio.Atualizar(f.ParaModelo())
}

func (serv FeiraServico) Deletar(ID int64) (*dominios.Feira, error) {
	f, err := serv.Repositorio.Recuperar(ID)

	if err != nil {
		return nil, err
	}

	return serv.Repositorio.Deletar(f.ParaModelo())
}

func (serv FeiraServico) RecuperarTudo() (*[]dominios.Feira, error) {
	return serv.Repositorio.RecuperarTudo()
}
