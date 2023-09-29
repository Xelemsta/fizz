FROM golang:alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/fizzbuzz ./cmd/fizz-buzz-api-server/main.go

EXPOSE 3001

ENTRYPOINT [ "/app/fizzbuzz", "--port", "3001" ]