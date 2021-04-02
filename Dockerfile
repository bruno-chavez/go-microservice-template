FROM golang:1.15 as builder

WORKDIR /service

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o microservice


FROM scratch

ENV PORT=8080

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /service/microservice .

EXPOSE ${PORT}
ENTRYPOINT ["/microservice"]
