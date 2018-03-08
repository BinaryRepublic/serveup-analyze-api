FROM golang:latest
WORKDIR /go/src/app
COPY ./src .
RUN mkdir ../soundfiles
RUN go get "github.com/gorilla/mux"
RUN go run main.go