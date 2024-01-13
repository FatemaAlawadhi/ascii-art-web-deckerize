# Use the official Go base image
FROM golang:latest

# Set labels for the image
LABEL maintainer="Fatema Awadhi - Asma Ahmed"
LABEL version="latest"
LABEL description="Ascii Art Generator"

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy all folders and files into the container
COPY . .

# Build the Go application (Image name is ascii-art-image)
RUN go build -o ascii-art

# Set the command to run when the container starts
CMD ["./ascii-art"]