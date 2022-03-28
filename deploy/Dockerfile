FROM golang:1.17-alpine AS builder

RUN apk update && apk upgrade && apk add --no-cache bash git openssh

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build GOOS=linux -o ./.bin/app ./cmd/app/main.go


FROM alpine

WORKDIR /app

COPY --from=builder ./app/.bin/ ./.bin/
COPY --from=builder ./app/config/ ./config/
COPY --from=builder ./app/docs/ ./docs/

EXPOSE 8080

CMD ["./app"]