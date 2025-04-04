FROM golang:1.23.1-alpine3.20

RUN apk add --no-cache \
    make

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make build 

EXPOSE 80

CMD [ "./bin/main" ]