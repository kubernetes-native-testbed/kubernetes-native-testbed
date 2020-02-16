FROM golang:1.13.6-buster as builder
COPY ./*.go ./
RUN go build -o ./app ./main.go
EXPOSE 8080
USER nobody
ENTRYPOINT ["./app"]
