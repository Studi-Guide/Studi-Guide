# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

ADD . /go/src/httpExample
WORKDIR /go/src/httpExample
RUN go get
RUN go install

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/go/bin/httpExample"]
