# Use an official Golang image as the build environment
FROM golang:1.21-alpine AS build

# Set the working directory inside the container
WORKDIR /log-app

# Copy go.mod and go.sum files first for dependency caching
COPY ./go.mod ./go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go app (replace `your-app` with your app name)
RUN go build -o loggen

# Use a minimal base image for the final build
FROM scratch

# Set working directory in final image
WORKDIR /root/app

# Copy the Go binary from the build stage
COPY --from=build /log-app/loggen loggen

# Command to run the
CMD ["./loggen"]
