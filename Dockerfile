FROM golang:1.18-alpine as builder

LABEL maintainer = "Yernur_chief"

WORKDIR /app

COPY . .
COPY 20221106152142_up.sql /app

RUN go build -o main ./main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app .

RUN apk add bash

ENTRYPOINT ["/app/main"]