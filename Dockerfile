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