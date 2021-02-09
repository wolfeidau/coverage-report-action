FROM golang:1.15 as builder

WORKDIR /src
COPY . .
ENV GO111MODULE=on

RUN go build -o /bin/action

FROM debian

RUN DEBIAN_FRONTEND=noninteractive apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /bin/action /bin/action
ENTRYPOINT ["/bin/action"]