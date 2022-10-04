FROM golang:1.19-alpine3.16 as builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app cmd/compose/websocket/todo/*.go

FROM alpine:3.16
COPY --from=builder /app ./
EXPOSE 14586
ENTRYPOINT ["./app"]
