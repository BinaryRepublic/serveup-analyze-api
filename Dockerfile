FROM golang:latest
WORKDIR /go/src
COPY ./src .
COPY config.toml ./config.toml
RUN mkdir ../soundfiles
RUN go get "github.com/gorilla/mux"
RUN go build main.go routes.go
ENV PATH=/go/src
CMD ["main"]