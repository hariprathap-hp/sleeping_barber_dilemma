# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY . .

# Command to run the executable
CMD ["go", "run", "main.go"]
