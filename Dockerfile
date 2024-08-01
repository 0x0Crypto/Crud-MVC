FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . ./

RUN CGO_ENABLED=1 go build -o mvc

EXPOSE 8080

CMD ["./mvc"]