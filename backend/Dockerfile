FROM golang:alpine

WORKDIR /go/src
COPY . .
RUN apk update
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o wwapp_be .

EXPOSE 8080
ENTRYPOINT ["/go/src/wwapp_be", "-env", "prod"]
