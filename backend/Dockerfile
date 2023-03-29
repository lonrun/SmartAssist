# Base image
FROM golang:1.16

# Set working directory
WORKDIR /app

# Copy necessary files
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the application
RUN go build -o app ./cmd/main.go

# Expose port 8080
EXPOSE 8080

# Start the application
CMD ["./app"]

# This Dockerfile defines the following steps:

# Use the golang:1.16 base image.
# Set the working directory to /app.
# Copy the go.mod and go.sum files to the container and run go mod download to download the dependencies.
# Copy all other files to the container.
# Build the application and name it app.
# Expose port 8080 to allow external traffic.
# Start the application by running ./app.