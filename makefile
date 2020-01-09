## Build all the examples
build:
	go build -o gorilla-example gorilla/gorilla.go
	go build -o gin-example gin/gin.go.

sanitize:
	go fmt ./...
	go vet ./...

## Start docker containers for Elasticsearch, Kibana and the APM server
start:
	docker-compose -f ./docker/docker-compose.yml up -d elasticsearch kibana apm mysql

## Stop all docker containers
stop:
	docker-compose -f ./docker/docker-compose.yml stop

## Stop and removes all docker containers
destroy:
	docker-compose -f ./docker/docker-compose.yml down -v --remove-orphans

## Clean up the .dat offset files
clean:
	rm -f ./gorilla-example
	rm -f ./gin-example
