FROM golang:alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY *.go ./
COPY db/migrations db/migrations

# Build
RUN go build -o main .

EXPOSE 8080

# Run
CMD ["./main"]
