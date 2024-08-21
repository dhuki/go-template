FROM golang:1.19-alpine

WORKDIR /app
COPY . ./

ENV CGO_ENABLED=0
RUN GRPC_HEALTH_PROBE_VERSION=v0.4.13 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

WORKDIR /app/cmd/grpc
ENTRYPOINT ["go", "run", "main.go"]
EXPOSE 8080