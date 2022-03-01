limpar:
	go clean

docker-iniciar:
	docker-compose up -d --force-recreate postgres	

docker-contruir:
	docker-compose build postgres

docker-finalizar:
	docker-compose down

compose-banco: docker-finalizar docker-contruir docker-iniciar 

compose-servico:
	docker build -t feira_registro . --target iniciar
	docker-compose build feira_servico
	docker-compose up -d feira_servico

swagger:
	swag init -g ./src/main.go

iniciar: compose-banco swagger
	docker build -t feira_registro . --target iniciar
	docker-compose build feira_servico
	docker-compose up -d feira_servico	

para:
	docker-compose down		