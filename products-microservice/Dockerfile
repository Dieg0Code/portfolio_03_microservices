FROM golang:1.22.1 AS builder

LABEL maintainer="github.com/dieg0code"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go .
COPY src/ src/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o products-microservice .

FROM alpine:3.19

RUN apk --no-cache add ca-certificates

RUN adduser -S -D -H -h /app appuser

USER appuser

WORKDIR /app

COPY --from=builder /app/products-microservice .

EXPOSE 8080

CMD ["./products-microservice"]