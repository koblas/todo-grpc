FROM golang:1.20-alpine as builder
WORKDIR /build
RUN apk --update add --no-cache ca-certificates openssl tzdata && update-ca-certificates
COPY go.mod go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app cmd/compose/publicapi/file/*.go

FROM scratch
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app ./
EXPOSE 14586
ENTRYPOINT ["./app"]
