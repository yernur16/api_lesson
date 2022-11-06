FROM golang:1.18-alpine as builder

LABEL maintainer = "Yernur_chief"

WORKDIR /app

COPY . .

RUN go build -o main ./main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app .

RUN apk add bash

EXPOSE 8080

ENTRYPOINT ["/app/main"]