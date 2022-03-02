package controladores

import (
	"net/http"

	"gitbuh.com/peevieira/feiras-api/src/servicos"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ArquivoControlador struct {
	Servico servicos.ArquivoServico
}

func (controle ArquivoControlador) Carregar() (r *ArquivoControlador) {
	controle.Servico = *servicos.ArquivoServico{}.Carregar()
	return &controle
}

// CarregarArquivos Tudo godoc
// @Summary Carregar dados das feiras a partir do CSV
// @Description Endpont que ira recuperar as informações das feiras de um arquivo CSV e salvar
// @Tags Arquivo CSV
// @Accept json
// @Produce json
// @Success 200 {array} dominios.Feira
// @Router /arquivo/ [get]
func (controle ArquivoControlador) CarregarArquivos(c *gin.Context) {
	feiraService := servicos.FeiraServico{}.Carregar()
	feiras := controle.Servico.ParaFeira()

	for _, feira := range *feiras {

		feira, err := feiraService.CriarObjetoPronto(&feira)

		if err != nil {
			logrus.Error("[ERRO] Feira não foi salva:", feira)
		}
	}

	c.JSON(http.StatusOK, feiras)
}
