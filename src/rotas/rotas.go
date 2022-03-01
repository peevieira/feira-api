package rotas

import (
	"gitbuh.com/peevieira/feiras-api/src/controladores"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ProcessarSolicitacao() {
	r := gin.New()
	r.Use(gin.Recovery())
	RegistrarRotasFeira(r)
	RegistrarRotasArquivo(r)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8181")
}

func RegistrarRotasFeira(r *gin.Engine) {
	g := r.Group("/api/v1/feira")

	g.POST("/", controladores.FeiraControlador{}.Carregar().Criar)
	g.PUT("/:id", controladores.FeiraControlador{}.Carregar().Atualizar)
	g.GET("/:id", controladores.FeiraControlador{}.Carregar().Recuperar)
	g.GET("/", controladores.FeiraControlador{}.Carregar().RecuperarTudo)
	g.GET("/filtro/:campo/:valor", controladores.FeiraControlador{}.Carregar().RecuperarComFiltro)
	g.DELETE("/:id", controladores.FeiraControlador{}.Carregar().Deletar)
}

func RegistrarRotasArquivo(r *gin.Engine) {
	g := r.Group("/api/v1/arquivo")

	g.GET("/", controladores.ArquivoControlador{}.Carregar().CarregarArquivos)
}
