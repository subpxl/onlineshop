# Build stage
FROM golang:1.17 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Runtime stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/app .

# Expose the port your application listens on (if needed)
EXPOSE 8000

# Define any environment variables (if needed)
# ENV ENV_VARIABLE_NAME value

# Run your Go application
CMD ["./app"]

