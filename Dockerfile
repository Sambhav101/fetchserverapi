# using the latest Go image
FROM golang:latest

# set current working directory as app
WORKDIR /app

# copy everything from the root directory
COPY . .

# install Go dependencies
RUN go mod download

# Build the Go application
RUN go build -o app .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./app"]
