#!/usr/bin/env bash

set -eo pipefail

if [ ! -d "./build" ]; then
  mkdir -p "./build/bin"
fi

if [ ! -f "./build/cosmovisor.tar.gz" ]; then
  wget "https://github.com/cosmos/cosmos-sdk/releases/download/cosmovisor%2Fv1.2.0/cosmovisor-v1.2.0-linux-amd64.tar.gz" -O "./build/cosmovisor.tar.gz"
fi

if [ ! -f "./build/bin/cosmovisor" ]; then
  (cd build && tar -zxvf "./cosmovisor.tar.gz" -C "./bin" && rm ./bin/*.md)
fi

cat <<EOF > "./build/Dockerfile"
# compile fx-core
FROM golang:1.18.2-alpine3.16 as builder

RUN apk add --no-cache git build-base linux-headers

WORKDIR /app

# download and cache go mod
COPY ./go.* ./
RUN go env -w GO111MODULE=on && go mod download

COPY . .

RUN make build

# build fx-core
FROM alpine:3.16

ENV DAEMON_HOME=/root/.fxcore
ENV DAEMON_NAME=fxcored

# optional
ENV DAEMON_ALLOW_DOWNLOAD_BINARIES=false
ENV DAEMON_RESTART_AFTER_UPGRADE=true
ENV DAEMON_RESTART_DELAY=1s
ENV DAEMON_POLL_INTERVAL=1s
ENV DAEMON_DATA_BACKUP_DIR=/root/.fxcore
ENV UNSAFE_SKIP_BACKUP=false
ENV DAEMON_PREUPGRADE_MAX_RETRIES=5

WORKDIR root

COPY ./build/bin/cosmovisor /usr/bin/cosmovisor
COPY --from=builder /app/build/bin/fxcored /usr/bin/fxcored

EXPOSE 26656/tcp 26657/tcp 26660/tcp 9090/tcp 1317/tcp 8545/tcp 8546/tcp

VOLUME ["/root"]

ENTRYPOINT ["cosmovisor"]
EOF

docker build --no-cache -f "./build/Dockerfile" -t functionx/fx-core-visor:latest .

# docker run --rm -v $HOME/.fxcore:/root/.fxcore functionx/fx-core-visor:latest init /usr/bin/fxcored
# docker run --name fxcore -v $HOME/.fxcore:/root/.fxcore functionx/fx-core-visor:latest run start