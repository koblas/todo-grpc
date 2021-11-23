FROM golang:1.17.1-alpine3.14 as builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app -gcflags=-G=3 cmd/core/oauth_user/*.go

FROM alpine:3.14
COPY --from=builder /app ./
EXPOSE 14586
ENTRYPOINT ["./app"]
