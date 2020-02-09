# Dockerfile References: https://docs.docker.com/engine/reference/builder/
# --------------------- IONIC Build ------------------------
FROM node:latest as ionicbuilder

COPY /front /www/app

RUN npm install -g ionic

WORKDIR /www/app
RUN npm install

RUN ionic build --engine=browser
RUN ls

# Start from the latest golang base image
FROM golang:latest

ADD back /go/src/httpExample
WORKDIR /go/src/httpExample
RUN go mod download
RUN go install ./cmd/...

WORKDIR /go/bin/ionic
COPY --from=ionicbuilder /www/app/www .

WORKDIR /go/bin

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/go/bin/studi-guide"]
