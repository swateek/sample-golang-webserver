ARG GOLANG_VERSION=1.17
ARG ALPINE_VERSION=3.14

# ================================================================
# Creating a build
# ================================================================

FROM golang:$GOLANG_VERSION AS builder

RUN apt-get update && apt-get install -y \
    build-essential

WORKDIR /go/src/build
COPY . .

RUN rm -f bin/* && go build -o bin/sample-webserver .

# ================================================================
# Creating an image with the above build
# ================================================================

#FROM alpine:$ALPINE_VERSION AS prod
FROM golang:$GOLANG_VERSION AS prod

ENV VAR_BIN_DIR="/var/bin/app"

RUN mkdir -p $VAR_BIN_DIR

COPY --from=builder /go/src/build/bin/sample-webserver $VAR_BIN_DIR/sample-webserver

WORKDIR $VAR_BIN_DIR
ENTRYPOINT ["./sample-webserver"]