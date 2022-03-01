package main

import (
	_ "gitbuh.com/peevieira/feiras-api/docs"

	"gitbuh.com/peevieira/feiras-api/src/banco"
	"gitbuh.com/peevieira/feiras-api/src/logs"
	"gitbuh.com/peevieira/feiras-api/src/rotas"
	"github.com/sirupsen/logrus"
)

// @title           Serviço de registro de Feiras API
// @version         1.0
// @description     Documentação do serviço de registro de Feiras

// @contact.name   Pedro Vieira
// @contact.url    http://www.swagger.io/support
// @contact.email  pedro.hvieira@live.com

// @host      localhost:8181
// @BasePath  /api/v1
func main() {

	deferArquivoLog := logs.InicializarArquivoDeLog()
	defer deferArquivoLog()
	logs.ConfigurarLogador()

	logrus.Info("Aplicação iniciada...")

	banco.ConectarComBancoDeDados()
	rotas.ProcessarSolicitacao()
}
