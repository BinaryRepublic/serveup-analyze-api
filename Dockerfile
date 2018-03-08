FROM golang:latest
WORKDIR /go/src/app
COPY ./app .
RUN go get "github.com/gorilla/mux"
RUN go run main.go