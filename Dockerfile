# Dockerfile References: https://docs.docker.com/engine/reference/builder/
# --------------------- IONIC Build ------------------------
FROM node:latest as ionicbuilder

ARG ionicproduction

COPY /front /www/app

RUN npm install -g @ionic/cli

WORKDIR /www/app
RUN npm install

RUN ionic build --engine=browser --localize ${ionicproduction}
RUN ls

# Start from the latest golang base image
FROM golang:latest as golangbuilder

COPY back /go/src/studi-guide
WORKDIR /go/src/studi-guide
RUN ls
RUN go mod download

# generate database schema/docs/mocks
RUN go generate ./...

# build go binaries
# Force the go compiler to use modules
RUN go build  -a -tags netgo -v  -ldflags '-w -extldflags "-static"' -o /go/bin ./cmd/...

# prepare db
RUN sh ./preparedb.sh
RUN cp ./db.sqlite3 ./../bin
WORKDIR /go/bin

FROM scratch

WORKDIR /go/bin/ionic
COPY --from=ionicbuilder /www/app/www .

WORKDIR /go/src
COPY --from=golangbuilder /go/src .

WORKDIR /go/bin
COPY --from=golangbuilder /go/bin .

# Expose port 8080 to the outside world
EXPOSE 8080
CMD ["./studi-guide"]
