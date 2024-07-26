FROM golang:1.22.5

#set destination for COPY
WORKDIR /app

# Download go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

RUN ls -la

# Build the application
RUN go build -o /urlshortner

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/urlshortner"]