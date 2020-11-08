FROM golang:alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

# RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/main.go

EXPOSE 8080

CMD [ "./main", "--port", "8080", "--host", "0.0.0.0" ]
