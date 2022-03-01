package logs

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func InicializarArquivoDeLog() func() {
	a, err := os.OpenFile("feira-api-logs", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		logrus.Fatalf("Erro ao tentar criar arquivo de log: %v", err)
	}

	escritor := io.MultiWriter(os.Stdout, a)
	logrus.SetOutput(escritor)
	return func() {
		defer a.Close()
	}
}

func ConfigurarLogador() {
	logrus.SetFormatter(&logrus.TextFormatter{})
}
