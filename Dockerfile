FROM golang:1.16

WORKDIR src/

COPY . .

RUN go build

EXPOSE 8080

CMD ./coder
