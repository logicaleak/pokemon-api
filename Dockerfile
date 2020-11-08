FROM golang:1.14

WORKDIR /app

COPY go.mod .
COPY go.sum .

# RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/main.go

EXPOSE 6390

CMD [ "./main", "--port", "6390" ]
