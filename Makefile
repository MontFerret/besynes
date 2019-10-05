export GOPATH
export GO111MODULE=on
export NODE_ENV=production
export CGO_ENABLED=1

export CC=gcc
export CGO_CFLAGS=-I/usr/local/include
export CGO_LDFLAGS=-Wl -I${ZMQ}/lib -I${GCC}/lib -lsodium -lzmq -lpthread -lstdc++ -lm -lc
export PKG_CONFIG_PATH=${ZMQ}/lib/pkgconfig

DIR_WORKER_SRC=./worker
DIR_APP_SRC=./app
DIR_BIN=./dist
DIR_PROTO=./proto
NODE_BIN=./node_modules/.bin
GO_ROOT=$(go env GOROOT)

default: build

start: generate
	qtdeploy -fast test desktop

build: generate
	qtdeploy -fast build desktop

install:
	go mod vendor && go mod tidy

generate:
	qtmoc desktop