# syntax=docker/dockerfile:1
FROM golang:1.18-alpine AS apigateway

WORKDIR /myapp

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build ./Api/cmd/main.go

CMD [ "./main" ]

FROM golang:1.18-alpine AS cartserver

WORKDIR /myapp

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build ./Servers/cart/main.go

CMD [ "./main" ]

FROM golang:1.18-alpine AS mongoclientserver

WORKDIR /myapp

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build ./Servers/mongoclient/main.go

CMD [ "./main" ]

FROM fluent/fluentd AS fluentd-elastic
RUN ["gem","install","fluent-plugin-elasticsearch","--no-rdoc","--no-ri"]