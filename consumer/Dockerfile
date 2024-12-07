# Use an official Golang image as the build environment
FROM golang:1.21-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o consumer

# Use a minimal base image for the final build
FROM alpine:latest

WORKDIR /root/
COPY --from=build /app/consumer .

CMD ["./consumer"]
