FROM golang:latest
WORKDIR /go/src/app
COPY ./app .
RUN go run main.go