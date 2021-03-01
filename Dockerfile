FROM golang:1.16-alpine AS build

ENV GO111MODULE=on

WORKDIR /go/src/app

LABEL maintainer="ashanaakh@gmail.com"

RUN apk add bash ca-certificates git gcc g++ libc-dev

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN export VERSION=$(cat VERSION) && \
    go build -ldflags "-X github.com/opencars/edrmvs/pkg/version.Version=$VERSION" -o /go/bin/extractor ./cmd/extractor/main.go && \
    go build -ldflags "-X github.com/opencars/edrmvs/pkg/version.Version=$VERSION" -o /go/bin/server ./cmd/server/main.go

FROM alpine

RUN apk update && apk upgrade && apk add curl

WORKDIR /app

COPY --from=build /go/bin/ ./
COPY ./config ./config

EXPOSE 8080

CMD ["./server"]
