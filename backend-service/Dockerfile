
# Use the official Go image as the base image
FROM golang:1.23-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o backend-service .

# Use a minimal alpine image for the final stage
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the binary from the build stage
COPY --from=build /app/backend-service /app/.env ./

# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["sh", "-c", "./backend-service up && ./backend-service start"]
