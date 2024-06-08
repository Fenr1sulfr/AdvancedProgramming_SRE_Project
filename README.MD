# URL Shortener

## Overview
This project is a URL shortening service implemented in Golang. It provides an API for creating and managing shortened URLs and includes features such as logging, metrics collection, and a health check.

## Features
- **URL Shortening**: Create and manage shortened URLs.
- **Metrics**: Collect and expose metrics using Prometheus.
- **Logging**: Structured logging with different levels based on the environment.
- **Health Checks**: Basic health checks and graceful shutdown handling.

## Requirements
- Golang 1.21 or later
- SQLite3
- Docker (optional, for containerized deployment)

## Getting Started

### Configuration
The configuration is loaded from a config file. Ensure the following environment variables are set:
- `ENV`: Environment (local, dev, prod)
- `ADDRESS`: Address to bind the server (e.g., `:8080`)
- `TIMEOUT`: Timeout duration for the server
- `IDLE_TIMEOUT`: Idle timeout duration for the server
- `HTTP_USER`: Basic auth username
- `HTTP_PASSWORD`: Basic auth password
- `SQLITE_PATH`: Path to the SQLite database file

### Running the Application

#### Locally
1. **Clone the repository**:
    ```sh
    git clone https://github.com/your-repo/url-shortener.git
    cd url-shortener
    ```

2. **Install dependencies**:
    ```sh
    go mod download
    ```

3. **Run the application**:
    ```sh
    go run main.go
    ```

#### Docker
1. **Build the Docker image**:
    ```sh
    docker build -t url-shortener .
    ```

2. **Run the Docker container**:
    ```sh
    docker run -p 8082:8082 url-shortener
    ```

### API Endpoints
- **Create URL**:
    ```
    POST /url
    Request Body:
    {
      "url": "https://example.com",
      "alias": "example"
    }
    Response:
    {
      "shortened_url": "http://localhost:8082/example"
    }
    ```

- **Redirect URL**:
    ```
    GET /{alias}
    Redirects to the original URL.
    ```

- **Metrics**:
    ```
    GET /url/metrics
    Exposes Prometheus metrics.
    ```

### Logging
Logs are written to `./logs/app.log`. The logging level is determined by the `ENV` variable:
- `local`: Pretty logs with error level
- `dev`: JSON logs with debug level
- `prod`: JSON logs with info level

### Graceful Shutdown
The server supports graceful shutdown and will handle system interrupts (`SIGINT`, `SIGTERM`). It allows up to 10 seconds for active connections to complete before shutting down.

## Dockerfile
The project includes a multi-stage Dockerfile:
- **Builder Stage**: Compiles the Go application.
- **Production Stage**: Uses an Alpine image for a minimal footprint.

```Dockerfile
# Stage 1: Building the application
FROM golang:1.21 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN apt-get update && apt-get install -y gcc g++ sqlite3
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o ./cmd/url-shortener ./cmd/url-shortener
RUN chmod a+x /app/cmd/url-shortener

# Stage 2: Production stage using Alpine
FROM alpine:latest
RUN apk update && apk --no-cache add ca-certificates sqlite
COPY --from=builder /app/cmd/url-shortener /url-shortener
RUN chmod a+x /url-shortener 
EXPOSE 8082
CMD ["./url-shortener"]
```

## Contributing
Contributions are welcome! Please open an issue or submit a pull request with your improvements.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.