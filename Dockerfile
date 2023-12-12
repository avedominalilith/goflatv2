RUN go build -o goflatv2 ./main.go
FROM golang:1.21-alpine

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o goflatv2 ./main.go

CMD ["./goflatv2"]