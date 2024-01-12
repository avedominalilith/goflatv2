FROM golang:1.21-alpine

WORKDIR /go/scr/goflatv2

COPY . .

RUN go build -o goflatv2

CMD ["./goflatv2"]

