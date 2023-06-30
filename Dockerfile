FROM golang:1.20-bullseye

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

RUN go build -v -o main .

CMD ["/app/main"]

