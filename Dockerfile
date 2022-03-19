FROM golang:1.17

EXPOSE 8888

WORKDIR /usr/local/grpc-mafia

COPY . /usr/local/grpc-mafia/

ENTRYPOINT go run server/*.go
