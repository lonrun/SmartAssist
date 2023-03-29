# Variables
BINARY_NAME=app
DOCKER_IMAGE_NAME=my-docker-image-name

# Go build command
build:
	go build -o $(BINARY_NAME) ./cmd/main.go

# Docker build command
docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

# Docker run command
docker-run:
	docker run -p 8080:8080 $(DOCKER_IMAGE_NAME)

# Clean up command
clean:
	rm -f $(BINARY_NAME)


# This Makefile defines four targets:

# build: Builds the binary file for the server application.
# docker-build: Builds the Docker image of the application.
# docker-run: Runs the application inside a Docker container.
# clean: Removes the generated binary file.
# You can use the make command to execute these targets. 
# For example, use make build to build the binary file for the server application.