#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
RUN apk add --no-cache --update go gcc g++
WORKDIR /app
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=1 GOOS=linux go build -o myapp cmd/server/main.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/myapp .
LABEL Name=shorturl Version=0.0.1
EXPOSE 4000
ENTRYPOINT ["/app/myapp"]
