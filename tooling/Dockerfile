FROM golang:1.15.2-alpine3.12 AS build
RUN apk --no-cache add bash g++ openssl git
RUN go get -u github.com/mitchellh/gox

ENV SWAGGER_VERSION=v0.25.0

ENV REPO="github.com/go-swagger/go-swagger/cmd/swagger"
RUN go get -u "${REPO}"
WORKDIR /go/src/github.com/go-swagger/go-swagger
RUN git checkout "${SWAGGER_VERSION}"
RUN gox --osarch "linux/amd64" -output /swagger \
      -ldflags "-s -w -X ${REPO}/commands.Version=${SWAGGER_VERSION} -X ${REPO}/commands.Commit=$(git rev-parse --short HEAD)" \
      "${REPO}"


FROM golang:1.15.2-alpine3.12
RUN apk add --update --no-cache make
COPY --from=build /swagger /usr/bin/swagger

ENV GOPATH=/go
WORKDIR /app
