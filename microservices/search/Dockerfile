FROM golang:1.13.6-buster as builder
COPY ./*.go ./
RUN go build -o /app ./main.go

FROM ubuntu:18.04
EXPOSE 8080
COPY --from=builder /app .
USER nobody
ENTRYPOINT ["./app"]


