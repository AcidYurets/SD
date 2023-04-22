FROM golang:1.20
WORKDIR /build

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

EXPOSE 4040

CMD ["./app"]

