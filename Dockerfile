FROM golang:latest

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]