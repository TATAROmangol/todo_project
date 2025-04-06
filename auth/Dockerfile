FROM golang:1.24.1-alpine3.20

RUN apk add --no-cache \
    make

WORKDIR /todo

COPY go.mod go.sum ./

RUN go mod download 

COPY . ./

RUN apk add --no-cache protobuf-dev
RUN make buildCmd 

CMD ["./bin/main"]


