FROM golang:1.23.1-alpine3.20

RUN apk add --no-cache \
    make

WORKDIR /auth

COPY .mod .sum ./

RUN go mod download 

COPY . ./

RUN make buildProto 

EXPOSE 50052

RUN make run


