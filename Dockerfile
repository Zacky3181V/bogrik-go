FROM registry.access.redhat.com/ubi9/go-toolset:latest AS builder

WORKDIR /opt/app-root/src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -o bogrik-go .

FROM registry.access.redhat.com/ubi9/ubi-minimal:latest

RUN microdnf install -y ca-certificates && \
    microdnf clean all

RUN useradd -u 1001 -r -g 0 -s /sbin/nologin \
    -c "Default Application User" default && \
    mkdir -p /app && \
    chown -R 1001:0 /app

WORKDIR /app

COPY --from=builder /opt/app-root/src/bogrik-go .

COPY --from=builder /opt/app-root/src/assets-compressed ./assets-compressed
COPY --from=builder /opt/app-root/src/templates ./templates

RUN chown -R 1001:0 /app && \
    chmod -R g=u /app

USER 1001

EXPOSE 8085

CMD ["./bogrik-go"]