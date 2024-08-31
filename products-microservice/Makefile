create_network:
	docker network create prod-network

start_test_db:
	docker run -d --name pg-prod --network prod-network -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgres -p 5432:5432 postgres:12.2

build:
	docker build -t prod-microservice .

run:
	docker run -d --name prod-microservice --network prod-network -p 8080:8080 prod-microservice

stop:
	docker stop pg-prod
	docker rm pg-prod
	docker stop prod-microservice
	docker rm prod-microservice
	docker network rm prod-network