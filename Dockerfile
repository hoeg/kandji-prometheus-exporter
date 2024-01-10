FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kandji-prometheus-exporter .

FROM scratch

WORKDIR /app

COPY --from=builder /app/kandji-prometheus-exporter .

EXPOSE 8080

CMD ["./kandji-prometheus-exporter"]
