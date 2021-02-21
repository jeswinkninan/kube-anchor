FROM golang:1.15.1 AS BUILDENV
WORKDIR /buildworkspace
ENV ARCH=linux
ENV CGO_ENABLED=0
COPY ./cmd/webhook-server/ ./
RUN go build -o kube-anchor-webhook-server

FROM alpine:3.12.3
COPY --from=BUILDENV /buildworkspace/kube-anchor-webhook-server /bin/kube-anchor-webhook-server
ENTRYPOINT ["kube-anchor-webhook-server"]
