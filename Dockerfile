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

COPY back /go/src/studi-guide
WORKDIR /go/src/studi-guide
RUN ls
RUN go mod download

# update swagger files
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag init -g ./cmd/studi-guide/main.go

# generate database schema
RUN go generate ./ent

# build go binaries
RUN go install ./cmd/...

WORKDIR /go/bin/ionic
COPY --from=ionicbuilder /www/app/www .
#RUN ls

WORKDIR /go/bin
RUN ls

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
# Use shell form to import rooms and then start server
CMD /go/bin/studi-guide-ctl migrate import rooms /go/src/studi-guide/rooms.json && /go/bin/studi-guide
