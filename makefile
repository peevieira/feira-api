limpar:
	go clean

docker-iniciar:
	docker-compose up -d --force-recreate postgres	

docker-construir:
	docker-compose build postgres

docker-finalizar:
	docker-compose down

compose-banco: docker-finalizar docker-construir docker-iniciar 

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

parar:
	docker-compose down	

logs:
	docker cp feiras-api_feira_servico_1:/src/feira-api-logs feira-api-logs
	



