FROM golang:1.20-alpine as builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app cmd/compose/workers/workers_user/*.go

FROM scratch
COPY --from=builder /app ./
ENTRYPOINT ["./app"]
