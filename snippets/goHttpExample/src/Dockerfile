# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest as builder

WORKDIR /go/src/httpExample
COPY . .
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s" -o ../../bin/httpExample .

FROM iron/base
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR /go
COPY --from=builder go ./


# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/go/bin/httpExample"]
