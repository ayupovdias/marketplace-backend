FROM golang:1.25-alpine

WORKDIR /app

RUN apk add --no-cache curl tar

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.3/migrate.linux-amd64.tar.gz \
    | tar xvz && mv migrate /usr/local/bin/migrate

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd

EXPOSE 8080

CMD ["sh", "-c", "sleep 20 && migrate -path migrations -database 'postgres://postgres:postgres@postgres:5432/market11?sslmode=disable' up && ./main"]