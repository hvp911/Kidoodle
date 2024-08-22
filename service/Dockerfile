FROM golang:1.19

WORKDIR /service

COPY . .

RUN go build -o content cmd/api/api.go

CMD ["./content"]
