FROM golang:1.13.6-buster as builder1
WORKDIR /app
COPY ./ /app/
RUN cd delivery-status; go get; go build -o /server

FROM busybox as builder2
RUN GRPC_HEALTH_PROBE_VERSION=v0.3.1 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

FROM ubuntu:18.04
EXPOSE 8080
COPY --from=builder1 /server .
COPY --from=builder2 /bin/grpc_health_probe /bin/grpc_health_probe
RUN chmod 755 ./server
USER nobody
ENTRYPOINT ["./server"]
