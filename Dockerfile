FROM golang:1.17-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git openssh

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build ./cmd/app

EXPOSE 8080

CMD ["./app"]