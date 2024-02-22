BUILD_PATH=build
BINARY_NAME=rinha
DOCKER_IMAGE_PREFIX=emanueljesusxavier1999
DOCKER_IMAGE_NAME=rinha-2024-q1

.PHONY: build exec

build:
	mkdir -p build
	go build -o build/$(BINARY_NAME) .

run:
	go run *.go

clean:
	rm -rf build/*

exec: clean build
	./$(BUILD_PATH)/$(BINARY_NAME)

docker-build-images:
	docker buildx build -t $(DOCKER_IMAGE_PREFIX)/$(DOCKER_IMAGE_NAME) .
	docker buildx build --platform linux/amd64 -t $(DOCKER_IMAGE_PREFIX)/$(DOCKER_IMAGE_NAME)-amd64 .

docker-push-images:
	docker push $(DOCKER_IMAGE_PREFIX)/$(DOCKER_IMAGE_NAME) 
	docker push $(DOCKER_IMAGE_PREFIX)/$(DOCKER_IMAGE_NAME)-amd64

build-docker: build-docker-images push-docker-images

docker-run:
	docker network create rinha
	docker run --name db -v ./db/script.sql:/docker-entrypoint-initdb.d/script.sql --network rinha -e POSTGRES_PASSWORD=1234 -d --restart=always --health-cmd="pg_isready -U postgres" --health-interval=10s --health-retries=5 --health-start-period=30s postgres
	docker run --name rinha --network rinha -p 9000:9000 -d $(DOCKER_IMAGE_PREFIX)/$(DOCKER_IMAGE_NAME)

start-docker:
	docker start db
	@sleep 10 && echo "Waiting database"
	docker start rinha

docker-stop:
	docker stop db
	@sleep 10 && echo "Waiting database"
	docker stop rinha

docker-clean:
	docker rm -f db rinha
	docker network rm rinha
