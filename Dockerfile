FROM golang:1.14-alpine as build
ARG TAG=production

WORKDIR /app

RUN apk update --no-cache \
  && apk add --no-cache \
    git \
    gcc \
    musl-dev

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o app main.go

RUN GO111MODULE=off go get -tags 'mysql' -u github.com/golang-migrate/migrate/cmd/migrate

FROM alpine:3.10

WORKDIR /app

RUN apk update --no-cache \
  && apk add --no-cache ca-certificates
RUN update-ca-certificates

COPY --from=build /app/app .

CMD ["./app"]