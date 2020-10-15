FROM golang:1.15

LABEL Mantainer="Isaac Rodríguez"

RUN mkdir /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

CMD ["/app/main"]
