
# it was last
# Stage 1: Build the application
FROM golang:1.20-alpine3.18 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files to the working directory
COPY go.mod go.sum ./

# Download and cache Go modules dependencies
RUN go mod download

# Copy the rest of the project files to the working directory
COPY . .


# Build the Go application
RUN go build -o main ./cmd/api/main.go

# Stage 2: Create the final production image
FROM alpine:3.14 as production

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage to the final image
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Set the entry point command for the container
CMD ["./main"]

# docker build -t my-golang-app .
# docker run -d --name my-golang-app-container my-golang-app
# docker exec -it my-golang-app-container /bin/sh
# docker rm my-golang-app-container

