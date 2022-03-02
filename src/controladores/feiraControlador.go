package controladores

import (
	"net/http"
	"strconv"

	"gitbuh.com/peevieira/feiras-api/src/modelos/dominios"
	"gitbuh.com/peevieira/feiras-api/src/servicos"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type FeiraControlador struct {
	Servico   servicos.FeiraServico
	Validador servicos.ValidadorServico
}

var (
	INI_PROC = ": Iniciando processamento da chamada"
	FIN_PROC = ": Finalizado processamento da chamada"
	ERR_PROC = ": Erro: "
)

func (controle FeiraControlador) Carregar() (r *FeiraControlador) {
	controle.Servico = *servicos.FeiraServico{}.Carregar()
	controle.Validador = *servicos.ValidadorServico{}.Carregar()
	return &controle
}

// Criar godoc
// @Summary Cria uma nova feira
// @Description Endpoint que permite criar uma nova feira ao enviar os dados da feira
// @Tags Feira
// @Accept json
// @Produce json
// @Param feira body dominios.Feira true "Body"
// @Success 200 {object} dominios.Feira
// @Router /feira/ [post]
func (controle FeiraControlador) Criar(c *gin.Context) {
	r := dominios.Feira{}

	if err := c.ShouldBindJSON(&r); err != nil {
		logrus.Error(c.Request.URL.String(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	v, err := controle.Validador.ValidarStructFeira(&r)

	if err != nil {
		logrus.Error(c.Request.URL.String(), err.Error())
		c.JSON(http.StatusBadRequest, v)
		return
	}

	f, err := controle.Servico.Criar(&r)

	if err != nil {
		logrus.Error(c.Request.URL.String(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, f)
}

// Recuperar godoc
// @Summary Recuperar uma feira
// @Description Endpont que permite passar o id, e retornar as informações da feira com esse id
// @Tags Feira
// @Accept json
// @Produce json
// @Param id path string true "Id da Feira"
// @Success 200 {object} dominios.Feira
// @Router /feira/:id [get]
func (controle FeiraControlador) Recuperar(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Params.ByName("id"), 10, 64)

	f, err := controle.Servico.Recuperar(id)

	if err != nil {
		logrus.Error(c.Request.URL.String(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, f)
}

// Recuperar Com Filtro godoc
// @Summary Recuperar feiras com filtros
// @Description Endpont que permite passar o tipo de filtro e o valor, e retornar as informações das feiras filtradas
// @Tags Feira
// @Accept json
// @Produce json
// @Param campo path string true "Campo do filtro" Enums(distrito, regiao5, nome, bairro)
// @Param valor path string true "Valor do campo do filtro"
// @Success 200 {array} dominios.Feira
// @Router /feira/filtro/:campo/:valor [get]
func (controle FeiraControlador) RecuperarComFiltro(c *gin.Context) {
	f, err := controle.Servico.RecuperarComFiltro(c.Params.ByName("campo"), c.Params.ByName("valor"))

	if err != nil {
		logrus.Error(c.Request.URL.String(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, f)
}

// Atualizar godoc
// @Summary Atualiza as informações de uma feira
// @Description Endpoint que permite atualizar uma feira ao enviar os dados da feira e também o id da feira
// @Tags Feira
// @Accept json
// @Produce json
// @Param feira body dominios.Feira true "Body"
// @Param id path string true "Id da Feira"
// @Success 200 {object} dominios.Feira
// @Router /feira/:id [put]
func (controle FeiraControlador) Atualizar(c *gin.Context) {
	r := dominios.Feira{}

	if err := c.ShouldBindJSON(&r); err != nil {
		logrus.Error(c.Request.URL.String(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	v, err := controle.Validador.ValidarStructFeira(&r)

	if err != nil {
		logrus.Error(c.Request.URL.String(), err.Error())
		c.JSON(http.StatusBadRequest, v)
		return
	}

	id, _ := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	f, err := controle.Servico.Atualizar(&r, id)

	if err != nil {
		logrus.Error(c.Request.URL.String(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, f)
}

// Deletar godoc
// @Summary Deletar as informações de uma feira
// @Description Endpoint que permite deletar uma feira ao enviar o id da feira
// @Tags Feira
// @Accept json
// @Produce json
// @Param id path string true "Id da Feira"
// @Success 200 {object} dominios.Feira
// @Router /feira/:id [delete]
func (controle FeiraControlador) Deletar(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Params.ByName("id"), 10, 64)

	_, err := controle.Servico.Deletar(id)

	if err != nil {
		logrus.Error(c.Request.URL.String(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"sucesso": "Registro deletado",
	})
}

// Recuperar Tudo godoc
// @Summary Recuperar todas as feiras
// @Description Endpont que vai retornar as informações de todas as feiras
// @Tags Feira
// @Accept json
// @Produce json
// @Success 200 {array} dominios.Feira
// @Router /feira/ [get]
func (controle FeiraControlador) RecuperarTudo(c *gin.Context) {
	f, err := controle.Servico.RecuperarTudo()

	if err != nil {
		logrus.Error(c.Request.URL.String(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, f)
}
