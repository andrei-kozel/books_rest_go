FROM golang:alpine

WORKDIR /app
COPY go mod go sum ./
run go mod download

COPY . .

RUN go build -o ./bin/api ./cmd/api \
    && go build -0 ./bin/migrate ./cmd/migrate

CMD ["/app/bin/api"]
EXPOSE 8080
