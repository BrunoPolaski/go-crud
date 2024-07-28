# Stage 1: Build the Go application
FROM golang:1.22 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules manifests
COPY go.mod go.sum ./

# Download dependencies, if go.sum is present it verifies checksums
RUN go mod download

# Copy the source code
COPY src src
COPY main.go main.go

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-crud .

# Stage 2: Create a lightweight runtime container
FROM golang:1.22-alpine AS runner

# Add a non-root user
RUN adduser -D bruneco

# Copy the built binary from the builder stage
COPY --from=builder /app/go-crud /app/go-crud

# Change ownership and permissions
RUN chown -R bruneco:bruneco /app
RUN chmod +x /app/go-crud

# Set the working directory
WORKDIR /app

# Expose the application port
EXPOSE 8080

# Switch to the non-root user
USER bruneco

# Command to run the application
CMD ["./go-crud"]
