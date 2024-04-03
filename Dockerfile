FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o app ./cmd/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/app /build/app
COPY configs/config.yaml /build/configs/config.yaml
CMD ["./app"]