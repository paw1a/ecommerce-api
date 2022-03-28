FROM golang:1.17-alpine AS builder

RUN apk update && apk upgrade && apk add --no-cache bash git openssh

WORKDIR /github.com/paw1a/ecommerce/
COPY . /github.com/paw1a/ecommerce/

RUN go mod download

RUN GOOS=linux go build -o ./.bin/app ./cmd/app/main.go


FROM alpine

WORKDIR /app

COPY --from=builder /github.com/paw1a/ecommerce/.bin/ .
COPY --from=builder /github.com/paw1a/ecommerce/config/ ./config/
COPY --from=builder /github.com/paw1a/ecommerce/docs/ ./docs/

EXPOSE 8080

CMD ["./app"]