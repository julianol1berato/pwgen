FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod .
COPY main.go .
COPY static/ static/

RUN go mod download
RUN go build -o main .

EXPOSE 80

CMD ["./main"]
