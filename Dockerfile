FROM golang:1.9.2
COPY main.go .
RUN go build -o app main.go
CMD ["./app"]
