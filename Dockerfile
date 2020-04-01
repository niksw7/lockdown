FROM golang:1.13
WORKDIR /app
COPY . .
RUN go build -o main .
CMD ["./main"]
