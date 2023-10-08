FROM golang:1.20-alpine

WORKDIR /app

COPY . ./

COPY config.yml.example config.yml

RUN go mod download && go mod verify \
	&& go build -v -o main . \
	&& rm -rf go.*

CMD ["/app/main"]
