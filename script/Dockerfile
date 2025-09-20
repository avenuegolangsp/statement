# nickfury
FROM golang:1.19-alpine

WORKDIR /app

COPY script/go.mod script/go.sum ./
RUN go mod download

COPY script/ .

CMD [ "go", "run", "main.go" ]
