# Download the latest version of Go using Alpine, a lightweight Linux version
FROM golang:alpine 


# Set the working directory inside the container
WORKDIR /groupie

# Copy all the files from your machine to the container
COPY . .

# Install Bash inside the container
RUN apk add bash 

# Build your Go program and name the output file "main"
RUN go build -o main .

# Expose port 8080 to allow access from outside the container
EXPOSE 8080

# Run the Go program when the container starts
CMD ["./main"]