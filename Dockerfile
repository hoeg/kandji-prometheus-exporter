FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o kandji-prometheus-exporter cmd/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/kandji-prometheus-exporter .

EXPOSE 8080

ENTRYPOINT [ "/app/kandji-prometheus-exporter" ]
