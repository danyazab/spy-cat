FROM golang:1.22.3-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

COPY ./pkg ./pkg
COPY ./cmd ./cmd
COPY ./internal ./internal

RUN go mod download

COPY . .

RUN ls -la /app

RUN go build -o spy-cat ./cmd

EXPOSE 8080

CMD ["./spy-cat"]