package repositorios

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"gitbuh.com/peevieira/feiras-api/src/modelos"
	"github.com/sirupsen/logrus"
)

type ArquivoRepositorio struct {
	Arquivo *os.File
	Err     error
}

func (repo ArquivoRepositorio) LerLinhas() *[]modelos.CampoArquivo {
	repo.Arquivo, repo.Err = os.Open("./arquivos/DEINFO_AB_FEIRASLIVRES_2014.csv")

	if repo.Err != nil {
		logrus.Error(repo.Err.Error())
	}

	defer repo.Arquivo.Close()

	linha1, err := bufio.NewReader(repo.Arquivo).ReadSlice('\n')

	if err != nil {
		logrus.Error(err.Error())
	}

	_, err = repo.Arquivo.Seek(int64(len(linha1)), io.SeekStart)

	if err != nil {
		logrus.Error(err.Error())
	}

	leitor := csv.NewReader(bufio.NewReader(repo.Arquivo))
	leitor.Comma = ','

	ret := []modelos.CampoArquivo{}

	logrus.Info("Inicio da importação das feiras do arquivo")

	for {
		linha, err := leitor.Read()

		if err != nil {
			logrus.Info("Fim da importação das feiras do arquivo")
			break
		}

		ret = append(ret, *repo.StringParaFeira(linha))

		if err == io.EOF {
			logrus.Info("Fim da importação das feiras do arquivo")
			break
		}
	}

	return &ret
}

func (repo ArquivoRepositorio) StringParaFeira(s []string) *modelos.CampoArquivo {
	return &modelos.CampoArquivo{
		Long:       ConverterParaFloat64(s[1]),
		Lat:        ConverterParaFloat64(s[2]),
		SetCens:    ConverterParaInt64(s[3]),
		AreaP:      ConverterParaFloat64(s[4]),
		CodDist:    ConverterParaInt64(s[5]),
		Distrito:   s[6],
		CodSubPref: ConverterParaInt64(s[7]),
		SubPrefe:   s[8],
		Regiao5:    s[9],
		Regiao8:    s[10],
		Nome_Feira: s[11],
		Registro:   s[12],
		Logradouro: s[13],
		Numero:     s[14],
		Bairro:     s[15],
		Referencia: s[16],
	}
}

func ConverterParaFloat64(n string) float64 {
	ret, err := strconv.ParseFloat(n, 64)

	if err != nil {
		logrus.Error(err.Error())
		return 0.
	}

	return ret
}

func ConverterParaInt64(n string) int64 {
	ret, err := strconv.ParseInt(n, 10, 64)

	if err != nil {
		logrus.Error(err.Error())
		return 0
	}

	return ret
}
