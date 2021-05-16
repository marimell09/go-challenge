FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/marimell09/stone-challenge/
WORKDIR /go/src/github.com/marimell09/stone-challenge
RUN go mod download
COPY . /go/src/github.com/marimell09/stone-challenge
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/stone-challenge github.com/marimell09/stone-challenge

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/marimell09/stone-challenge/build/stone-challenge /usr/bin/mariana
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/mariana"]

