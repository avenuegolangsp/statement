# processor
FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /processor cmd/processor/main.go

EXPOSE 8080

CMD [ "/processor" ]