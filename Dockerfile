FROM golang:1.20.3-alpine

# Set the working directy to /app
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod and go.sum ./

# Download and install any required Go dependencies
RUN go mod download

# Copy the entrie source code to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port specified by the PORT environment variable
EXPOSE 3000

# Set the entry point of the container to the executable
CMD ["./main"]