FROM golang:1.24.1-alpine3.20

RUN apk add --no-cache \
    make

WORKDIR /todo

COPY go.mod go.sum ./

RUN go mod download 

COPY . ./

RUN make buildProto 

EXPOSE 50052

RUN make run


