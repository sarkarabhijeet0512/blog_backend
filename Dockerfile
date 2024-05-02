FROM golang:1.20.13-alpine3.19 AS builder
# Setup base software for building app
RUN apk update && \
    apk add bash ca-certificates git gcc g++ libc-dev binutils file librdkafka-dev pkgconf libc6-compat 

WORKDIR /srv

# Download dependencies
ADD go.mod go.sum ./
RUN go mod download && go mod verify

# Copy an application's source
COPY ./ ./

# Go install will create multiple binary files in GOBIN path. 
# `cmd` dir in source code contains multiple directories. All these directories are `main` packages.
# This is a multi-main package pattern where multiple binary files can be created from a single source code.  
# Binary file names will be as per the directory names present in cmd dir. 
# Mark the build as statically linked.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -tags musl \
    -installsuffix 'static' \
    -o /blog_api \
    ./cmd/...

# Prepare executor image
FROM alpine:3.19.0 AS production

RUN apk update && \
    apk add ca-certificates libc6-compat && \
    rm -rf /var/cache/apk/*

# Add timezone data to image as its not bundled in project binary.
RUN apk add --no-cache tzdata

WORKDIR /srv

# Import executable from builder stage.
COPY --from=builder /blog_api /blog_api
COPY --from=builder /srv .

CMD ["/blog_api"]
