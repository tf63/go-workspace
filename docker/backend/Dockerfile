# syntax=docker/dockerfile:1
FROM golang:1.19-alpine

# ワークディレクトリの指定
WORKDIR /app

# パッケージのインストール
RUN apk update --no-cache \
    && apk add --no-cache \
    git \
    gcc \
    musl-dev \
    make \
    wget

# モジュールのダウンロード
COPY go.mod go.sum Makefile ./
RUN go mod download
RUN make install-tools
