make create_network:
	docker network create sales-network


start_test_db:
	docker run -d --name dynamodb-sales -p 8000:8000 --network sales-network amazon/dynamodb-local

add_test_table:
	aws dynamodb create-table \
	--table-name Sales \
	--attribute-definitions \
		AttributeName=saleID,AttributeType=S \
	--key-schema \
		AttributeName=saleID,KeyType=HASH \
	--provisioned-throughput \
		ReadCapacityUnits=5,WriteCapacityUnits=5 \
	--endpoint-url http://localhost:8000

stop_test_db:
	docker stop dynamodb-sales
	docker rm dynamodb-sales

build:
	docker build --no-cache -t sales-microservice .

run:
	docker run --name sales-microservice -p 8082:8082 --network sales-network -v ~/.aws:/root/.aws -e AWS_REGION=sa-east-1 -e AWS_SDK_LOAD_CONFIG=1 sales-microservice


stop:
	docker stop sales-microservice
	docker rm sales-microservice
	docker stop dynamodb-sales
	docker rm dynamodb-sales
	docker network rm sales-network