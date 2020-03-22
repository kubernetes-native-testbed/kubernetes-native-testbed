FROM ubuntu:18.04 as builder1
RUN apt-get update && apt-get install curl gnupg git pkg-config -y && apt-get clean && rm -rf /var/lib/apt/lists/*
ENV GOPATH=/go
ENV GOBIN=/go/bin
ENV GO_VERSION=1.13.6
RUN mkdir -p $GOPATH
RUN curl -LO https://dl.google.com/go/go$GO_VERSION.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz && \
    rm go$GO_VERSION.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin

RUN curl -s https://packages.confluent.io/deb/5.4/archive.key | apt-key add -
RUN apt-get update && apt-get -y install software-properties-common && \
    add-apt-repository "deb [arch=amd64] https://packages.confluent.io/deb/5.4 stable main" && \
    apt-get update && apt-get -y install librdkafka-dev librdkafka1 librdkafka++1 && apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY ./ /app/
RUN cd point; go get; go build -o /server

FROM busybox as builder2
RUN GRPC_HEALTH_PROBE_VERSION=v0.3.1 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

FROM ubuntu:18.04
EXPOSE 8080
RUN apt-get update && apt-get install curl gnupg -y && apt-get clean && rm -rf /var/lib/apt/lists/*
RUN curl -s https://packages.confluent.io/deb/5.4/archive.key | apt-key add -
RUN apt-get update && apt-get -y install software-properties-common && \
    add-apt-repository "deb [arch=amd64] https://packages.confluent.io/deb/5.4 stable main" && \
    apt-get update && apt-get -y install librdkafka-dev librdkafka1 librdkafka++1 && apt-get clean && rm -rf /var/lib/apt/lists/*

COPY --from=builder1 /server .
COPY --from=builder2 /bin/grpc_health_probe /bin/grpc_health_probe
RUN chmod 755 ./server
USER nobody
ENTRYPOINT ["./server"]



