# backend/Dockerfile
FROM golang:1.20-alpine

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY config/ config/
COPY controllers/ controllers/
COPY middlewares/ middlewares/
COPY models/ models/
COPY routes/ routes/
COPY utils/ utils/
COPY main.go ./

# Build the Go app
RUN go build -o main .

# Expose port
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
