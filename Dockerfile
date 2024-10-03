FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Install the air tool for live reload
RUN go install github.com/air-verse/air@latest

# Copy go.mod, go.sum and all into the container
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the project files including the .air.toml and source code
COPY . .

# Run the application using the air config
CMD ["sh", "-c", "sleep 30 && air -c .air.toml"]
