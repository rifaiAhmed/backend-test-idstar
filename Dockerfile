FROM golang:1.24.1

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod tidy

COPY . .

COPY .env .

RUN go build -o backend-test

RUN chmod +x backend-test

EXPOSE 8080

CMD ["./backend-test"]