FROM golang:1.23-alpine AS build

ENV GO111MODULE=on

WORKDIR /go/src/app

RUN apk add bash ca-certificates git gcc g++ libc-dev curl make

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN export BLDDIR=/go/bin && \
    make clean && \
    make build

FROM alpine

RUN apk update && apk upgrade && apk add curl

WORKDIR /app

COPY --from=build /go/bin/ ./
COPY ./config ./config

EXPOSE 8080 3000

CMD ["./server"]
