# feira-api

Aplicação REST API que fornece a manutenção de Feiras 

### Contruído com

* [GoLang](https://golang.org/)
* [Gin](https://github.com/gin-gonic/gin)
* [Gorm](https://gorm.io/index.html)
* [Validator](https://github.com/go-playground/validator)
* [Swag](https://github.com/swaggo/swag)
* [Logrus](https://github.com/sirupsen/logrus)

### Instalação

1. Clonar o repositório
   ```sh
   git clone https://github.com/peevieira/feira-api.git
   ```
## Uso

1. Rodar serviço e Banco como container Docker
   ```sh
   make iniciar
   ```
2. Rodar serviço como container Docker
   ```sh
   make compose-servico
   ```   
3. Rodar Banco como container Docker
   ```sh
   make compose-banco
   ```    
4. Gerar documentação da API
   ```sh
   make swagger
   ```     
### Importar CSV

```sh 
    curl --location --request GET 'http://localhost:8181/api/v1/arquivo/'
```
   
### Criar

```sh 
    curl --location --request POST 'http://localhost:8181/api/v1/feira' \
    --header 'Content-Type: application/json' \
    --data-raw '{
      "census_sector": "string",
      "DistritoNome": "string",
      "SubPrefeituraNome": "string",
      "Longitude": 1.0,
      "Latitude": 1.0,
      "Censo": 1,
      "Area": 1.0,
      "Regiao5": "string",
      "Regiao8": "string",
      "Nome": "string",
      "Registro": "string",
      "Logradouro": "string",
      "Numero": "string",
      "Bairro": "string",
      "Referencia": "string"
    }'  
```

### Recuperar

```sh 
    curl --location --request GET 'http://localhost:8181/api/v1/feira/1'
```

### Deletar

```sh 
    curl --location --request DELETE 'http://localhost:8181/api/v1/feira/1'
```

### Atualizar

```sh 
    curl --location --request PUT 'http://localhost:8181/api/v1/feira/1' \
    --header 'Content-Type: application/json' \
    --data-raw '{
      "census_sector": "string",
      "DistritoNome": "string",
      "SubPrefeituraNome": "string",
      "Longitude": 1.0,
      "Latitude": 1.0,
      "Censo": 1,
      "Area": 1.0,
      "Regiao5": "string",
      "Regiao8": "string",
      "Nome": "string",
      "Registro": "string",
      "Logradouro": "string",
      "Numero": "string",
      "Bairro": "string",
      "Referencia": "string"
    }'
```

### Procurar
```sh 
    curl --location --request GET 'http://localhost:8181/v1/feira/filtro/nome/string'
```
# Documentação

[Documentação](http://localhost:8181/docs/index.html) {http://localhost:8181/docs/index.html}
