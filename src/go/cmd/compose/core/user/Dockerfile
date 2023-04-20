FROM golang:1.19-alpine3.16 as builder
WORKDIR /build
ADD go.mod go.sum /build/
RUN go mod download
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app cmd/compose/core/user/*.go

FROM alpine:3.16
COPY --from=builder /app ./
EXPOSE 14586
CMD ["./app"]
