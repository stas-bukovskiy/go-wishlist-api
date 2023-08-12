# build stage
FROM golang:1.20.2 AS build-env

WORKDIR /go/src/app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies.
# Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/user-service ./cmd/main.go

# run stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=build-env /go/bin/user-service .
COPY --from=build-env /go/src/app/.env .
COPY --from=build-env /go/src/app/configs/ ./configs/

# Expose port 8080 to the outside world
EXPOSE 8000

#Command to run the executable
CMD ["./user-service"]