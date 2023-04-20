#download dependency stage
FROM golang:alpine AS builder
ENV GO111MODULE=on
RUN apk add --no-cache git make bash
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o /go/bin/app ./main.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN ln -sf /proc/1/fd/1 /tmp/application.log
COPY --from=builder /go/bin/app /app
ENTRYPOINT /app
LABEL Name=devcode-golang Version=0.0.1
EXPOSE 3030
