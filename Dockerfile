FROM golang:alpine

WORKDIR /app

COPY . .

RUN go get

RUN go bulid

EXPOSE 8000

CMD ["./main"]

