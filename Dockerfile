FROM golang:latest
WORKDIR /go/src
COPY ./src .
COPY config.toml ./config.toml
RUN mkdir ../soundfiles
RUN go get "github.com/gorilla/mux"
RUN go get "github.com/BurntSushi/toml"
RUN go build main.go routes.go
ENV GOPATH=/go/src
CMD ["/go/src/main"]