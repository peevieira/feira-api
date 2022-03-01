package servicos

import (
	"gitbuh.com/peevieira/feiras-api/src/modelos"
	"gitbuh.com/peevieira/feiras-api/src/modelos/dominios"
	"github.com/go-playground/validator/v10"
)

type ValidadorServico struct {
	Validador *validator.Validate
}

func (serv ValidadorServico) Carregar() *ValidadorServico {
	serv.Validador = validator.New()
	return &serv
}

func (serv ValidadorServico) ValidarStructFeira(f *dominios.Feira) (v *[]modelos.ValidadorErro, err error) {
	err = serv.Validador.Struct(f)

	if err == nil {
		return nil, nil
	}

	r := []modelos.ValidadorErro{}

	for _, err := range err.(validator.ValidationErrors) {

		err.Value()

		r = append(r, modelos.ValidadorErro{
			Namespace: err.Namespace(),
			Field:     err.Field(),
			Tag:       err.Tag(),
			Type:      err.Type().String(),
		})
	}

	return &r, err
}
