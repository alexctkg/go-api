FROM golang:1.14.7-alpine


ENV GIN_MODE=release

WORKDIR $GOPATH/src/tdez

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]