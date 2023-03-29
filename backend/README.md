# Project Name

This is a backend server application that provides APIs for PDF summary, ticket booking and device control. It is built in Go and utilizes various third-party packages to provide fast and efficient performance.

## Features

### Read PDFs and give a summary

- Endpoint: `/pdf`
- Method: `POST`
- Request body: PDF file (multipart/form-data)
- Response body: Summary of the PDF

### Book tickets for airlines and trains

- Endpoint: `/ticket`
- Method: `POST`
- Request body: Ticket information (JSON)
- Response body: Confirmation of the booking

### Control your home devices

- Endpoint: `/device`
- Method: `POST`
- Request body: Device information (JSON)
- Response body: Confirmation of the device control

## Technologies Used

- Go
- Third-party libraries
- Docker
- Docker Compose

## Getting Started

### Prerequisites

- Go
- Docker
- Docker Compose

### Installation

1. Clone the repository
2. Build the Docker image: `docker build -t <image-name> .`
3. Start the Docker container: `docker run -p <host-port>:<container-port> <image-name>`
4. Access the APIs through `localhost:<host-port>`

## Contributors

- [Contributor 1](https://github.com/contributor-1)
- [Contributor 2](https://github.com/contributor-2)

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).
