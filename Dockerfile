# Load from docker
FROM golang:latest

# Define url target
ENV REDIRECT_TARGET https://google.com

# Copy files to workdir
COPY . /go/src/app
WORKDIR /go/src/app

# Load libraries
RUN go get ./

# Build application
RUN go build -o app

# Start application
CMD ./app

# Set port
EXPOSE 80