FROM golang:1.23.2
# set working directory
WORKDIR /app

COPY . .
# Download and install the dependencies
RUN go mod tidy

# Copy the source code

#EXPOSE the port
EXPOSE 8080

# Run the executable
CMD ["go", "run", "cmd/api-gateway/main.go"]