FROM golang:1.13
WORKDIR /app
COPY . .
RUN go build -o main .
RUN  ls | grep -v main$ | xargs rm -rf
CMD ["./main"]
