package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"gitbuh.com/peevieira/feiras-api/src/banco"
	"gitbuh.com/peevieira/feiras-api/src/controladores"
	"gitbuh.com/peevieira/feiras-api/src/modelos/dominios"
	"gitbuh.com/peevieira/feiras-api/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() (*gin.Engine, *mocks.FeiraMock) {
	banco.ConectarComBancoDeDados()
	mock := mocks.FeiraMock{DB: banco.DB}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	return r, &mock
}

func TestListarTodosAsFeiras(t *testing.T) {

	r, _ := SetupDasRotasDeTeste()

	r.GET("/", controladores.FeiraControlador{}.Carregar().RecuperarTudo)

	req, _ := http.NewRequest("GET", "/", nil)

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, "Status HTTP precisa ser 200")
}

func TestIncluirUmaFeira(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	r.POST("/", controladores.FeiraControlador{}.Carregar().Criar)

	feiraJSON, _ := json.Marshal(mock.CriarDominioFeira())

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(feiraJSON))

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, "Status HTTP precisa ser 200")

	var feiraRet dominios.Feira
	json.Unmarshal(resposta.Body.Bytes(), &feiraRet)

	assert.Equal(t, mock.IdRegistroTest(), feiraRet.ID, "Inclusão da feira falhou")
}

func TestAlterarUmaFeira(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	r.PUT("/:id", controladores.FeiraControlador{}.Carregar().Atualizar)

	feira := mock.RecuperarUmaFeiraDoBanco()
	feira.Bairro = "Bairro alterado"
	feiraJSON, _ := json.Marshal(feira)

	req, _ := http.NewRequest("PUT", fmt.Sprintf("/%d", mock.IdRegistroTest()), bytes.NewBuffer(feiraJSON))

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, "Status HTTP precisa ser 200")

	var feiraRet dominios.Feira
	json.Unmarshal(resposta.Body.Bytes(), &feiraRet)

	assert.Equal(t, "Bairro alterado", feiraRet.Bairro, "Atualização da feira falhou")
}

func TestRecuperarUmaFeira(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	r.GET("/:id", controladores.FeiraControlador{}.Carregar().Recuperar)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/%d", mock.IdRegistroTest()), nil)

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, "Status HTTP precisa ser 200")

	var feiraRet dominios.Feira
	json.Unmarshal(resposta.Body.Bytes(), &feiraRet)

	assert.Equal(t, "Feira mock de teste", feiraRet.Nome, "Recuperação da feira falhou")
	assert.Equal(t, mock.IdRegistroTest(), feiraRet.ID, "Recuperação da feira falhou")
}

func TestDeletarUmaFeira(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	r.DELETE("/:id", controladores.FeiraControlador{}.Carregar().Deletar)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/%d", mock.IdRegistroTest()), nil)

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, "Status HTTP precisa ser 200")

	retorno, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mock.RespostasAoDeletar(), string(retorno), "Delete da feira falhou")
}

func TestFiltrarUmaFeiraRegiao5(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	feira := mock.CriarUmaFeiraParaFiltrarRegiao5()

	r.GET("/filtro/:campo/:valor", controladores.FeiraControlador{}.Carregar().RecuperarComFiltro)

	req, _ := http.NewRequest("GET", "/filtro/regiao5/filtro", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var feiraRet dominios.Feira
	json.Unmarshal(resposta.Body.Bytes(), &feiraRet)
	assert.Equal(t, feira.ID, feiraRet.ID, "Filtro por regiao5 não está funcionando")
}

func TestFiltrarUmaFeiraNome(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	feira := mock.CriarUmaFeiraParaFiltrarNome()

	r.GET("/filtro/:campo/:valor", controladores.FeiraControlador{}.Carregar().RecuperarComFiltro)

	req, _ := http.NewRequest("GET", "/filtro/nome/filtro", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var feiraRet dominios.Feira
	json.Unmarshal(resposta.Body.Bytes(), &feiraRet)
	assert.Equal(t, feira.ID, feiraRet.ID, "Filtro por nome não está funcionando")
}

func TestFiltrarUmaFeiraBairro(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	feira := mock.CriarUmaFeiraParaFiltrarBairro()

	r.GET("/filtro/:campo/:valor", controladores.FeiraControlador{}.Carregar().RecuperarComFiltro)

	req, _ := http.NewRequest("GET", "/filtro/bairro/filtro", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var feiraRet dominios.Feira
	json.Unmarshal(resposta.Body.Bytes(), &feiraRet)
	assert.Equal(t, feira.ID, feiraRet.ID, "Filtro por bairro não está funcionando")
}

func TestFiltrarUmaFeiraDistrito(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	feira := mock.CriarUmaFeiraParaFiltrarDistrito()

	r.GET("/filtro/:campo/:valor", controladores.FeiraControlador{}.Carregar().RecuperarComFiltro)

	req, _ := http.NewRequest("GET", "/filtro/distrito/filtro", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var feiraRet dominios.Feira
	json.Unmarshal(resposta.Body.Bytes(), &feiraRet)
	assert.Equal(t, feira.ID, feiraRet.ID, "Filtro por distrito não está funcionando")
}

func TestDeletarFeiraFiltradaRegiao5(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	r.DELETE("/:id", controladores.FeiraControlador{}.Carregar().Deletar)

	req, _ := http.NewRequest("DELETE", "/-10", nil)

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, "Status HTTP precisa ser 200")

	retorno, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mock.RespostasAoDeletar(), string(retorno), "Delete da feira falhou")
}

func TestDeletarFeiraFiltradaNome(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	r.DELETE("/:id", controladores.FeiraControlador{}.Carregar().Deletar)

	req, _ := http.NewRequest("DELETE", "/-20", nil)

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, "Status HTTP precisa ser 200")

	retorno, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mock.RespostasAoDeletar(), string(retorno), "Delete da feira falhou")
}

func TestDeletarFeiraFiltradaBairro(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	r.DELETE("/:id", controladores.FeiraControlador{}.Carregar().Deletar)

	req, _ := http.NewRequest("DELETE", "/-30", nil)

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, "Status HTTP precisa ser 200")

	retorno, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mock.RespostasAoDeletar(), string(retorno), "Delete da feira falhou")
}

func TestDeletarFeiraFiltradaDistrito(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	r.DELETE("/:id", controladores.FeiraControlador{}.Carregar().Deletar)

	req, _ := http.NewRequest("DELETE", "/-40", nil)

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, "Status HTTP precisa ser 200")

	retorno, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mock.RespostasAoDeletar(), string(retorno), "Delete da feira falhou")
}

func TestObrigatoriedadeCampoDistritoNome(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	r.POST("/", controladores.FeiraControlador{}.Carregar().Criar)

	feiraJSON, _ := json.Marshal(mock.CriarDominioFeiraSemValor("dist"))

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(feiraJSON))

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	retorno, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mock.RespostaValidador("dist"), string(retorno), "Validador do campo DistritoNome não está funcionando")
}

func TestObrigatoriedadeCampoSubPrefeituraNome(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	r.POST("/", controladores.FeiraControlador{}.Carregar().Criar)

	feiraJSON, _ := json.Marshal(mock.CriarDominioFeiraSemValor("pref"))

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(feiraJSON))

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	retorno, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mock.RespostaValidador("pref"), string(retorno), "Validador do campo SubPrefeituraNome não está funcionando")
}

func TestObrigatoriedadeCampoLongitude(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	r.POST("/", controladores.FeiraControlador{}.Carregar().Criar)

	feiraJSON, _ := json.Marshal(mock.CriarDominioFeiraSemValor("long"))

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(feiraJSON))

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	retorno, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mock.RespostaValidador("long"), string(retorno), "Validador do campo Longitude não está funcionando")
}

func TestObrigatoriedadeCampolatitude(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	r.POST("/", controladores.FeiraControlador{}.Carregar().Criar)

	feiraJSON, _ := json.Marshal(mock.CriarDominioFeiraSemValor("lat"))

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(feiraJSON))

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	retorno, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mock.RespostaValidador("lat"), string(retorno), "Validador do campo Latitude não está funcionando")
}

func TestObrigatoriedadeCampoCenso(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	r.POST("/", controladores.FeiraControlador{}.Carregar().Criar)

	feiraJSON, _ := json.Marshal(mock.CriarDominioFeiraSemValor("censo"))

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(feiraJSON))

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	retorno, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mock.RespostaValidador("censo"), string(retorno), "Validador do campo Censo não está funcionando")
}

func TestObrigatoriedadeCampoArea(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	r.POST("/", controladores.FeiraControlador{}.Carregar().Criar)

	feiraJSON, _ := json.Marshal(mock.CriarDominioFeiraSemValor("area"))

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(feiraJSON))

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	retorno, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mock.RespostaValidador("area"), string(retorno), "Validador do campo Area não está funcionando")
}

func TestObrigatoriedadeCampoRegiao5(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	r.POST("/", controladores.FeiraControlador{}.Carregar().Criar)

	feiraJSON, _ := json.Marshal(mock.CriarDominioFeiraSemValor("reg5"))

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(feiraJSON))

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	retorno, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mock.RespostaValidador("reg5"), string(retorno), "Validador do campo Regiao5 não está funcionando")
}

func TestObrigatoriedadeCampoRegiao8(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	r.POST("/", controladores.FeiraControlador{}.Carregar().Criar)

	feiraJSON, _ := json.Marshal(mock.CriarDominioFeiraSemValor("reg8"))

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(feiraJSON))

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	retorno, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mock.RespostaValidador("reg8"), string(retorno), "Validador do campo Regiao8 não está funcionando")
}

func TestObrigatoriedadeCampoNome(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	r.POST("/", controladores.FeiraControlador{}.Carregar().Criar)

	feiraJSON, _ := json.Marshal(mock.CriarDominioFeiraSemValor("nome"))

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(feiraJSON))

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	retorno, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mock.RespostaValidador("nome"), string(retorno), "Validador do campo Nome não está funcionando")
}

func TestObrigatoriedadeCampoRegistro(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	r.POST("/", controladores.FeiraControlador{}.Carregar().Criar)

	feiraJSON, _ := json.Marshal(mock.CriarDominioFeiraSemValor("registro"))

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(feiraJSON))

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	retorno, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mock.RespostaValidador("registro"), string(retorno), "Validador do campo Registro não está funcionando")
}

func TestObrigatoriedadeCampoLogradouro(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	r.POST("/", controladores.FeiraControlador{}.Carregar().Criar)

	feiraJSON, _ := json.Marshal(mock.CriarDominioFeiraSemValor("logra"))

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(feiraJSON))

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	retorno, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mock.RespostaValidador("logra"), string(retorno), "Validador do campo Logradouro não está funcionando")
}

func TestObrigatoriedadeCampoNumero(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	r.POST("/", controladores.FeiraControlador{}.Carregar().Criar)

	feiraJSON, _ := json.Marshal(mock.CriarDominioFeiraSemValor("num"))

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(feiraJSON))

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	retorno, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mock.RespostaValidador("num"), string(retorno), "Validador do campo Número não está funcionando")
}

func TestObrigatoriedadeCampoBairro(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	r.POST("/", controladores.FeiraControlador{}.Carregar().Criar)

	feiraJSON, _ := json.Marshal(mock.CriarDominioFeiraSemValor("bairro"))

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(feiraJSON))

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	retorno, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mock.RespostaValidador("bairro"), string(retorno), "Validador do Bairro nome não está funcionando")
}

func TestObrigatoriedadeCampoReferencia(t *testing.T) {
	r, mock := SetupDasRotasDeTeste()

	r.POST("/", controladores.FeiraControlador{}.Carregar().Criar)

	feiraJSON, _ := json.Marshal(mock.CriarDominioFeiraSemValor("ref"))

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(feiraJSON))

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	retorno, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mock.RespostaValidador("ref"), string(retorno), "Validador do campo Regerencia não está funcionando")
}
