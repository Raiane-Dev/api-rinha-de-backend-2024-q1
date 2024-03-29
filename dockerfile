FROM golang:1.21.1-alpine AS build

LABEL project="API da rinha - by Raiane-Dev"
LABEL author="raiane.dev@gmail.com"

RUN apk update && apk add --no-cache gcc musl-dev

## CONFIG GOLANG
ENV PATH="$PATH:$(go env GOPATH)/bin"
ENV CGO_ENABLED 1
ENV GOPATH /go
ENV GOCACHE /go-build
ENV GOOS linux
ENV GOARCH amd64

WORKDIR /go/src

COPY ./backend .

RUN go mod tidy

RUN cd cmd && \
		go build -o /bin/main


FROM alpine:latest as finally

WORKDIR /usr/local/bin/app

COPY --from=build /bin/main .
COPY ./database/schemas /data/schemas

EXPOSE 443/tcp

ENTRYPOINT ["/usr/local/bin/app/main"]