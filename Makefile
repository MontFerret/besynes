export GOPATH
export GO111MODULE=on
export NODE_ENV=production
export CGO_ENABLED=1

ZMQ?=/usr/local/Cellar/zeromq/4.3.2
GCC?=/usr/local/Cellar/gcc/9.2.0

export CC=gcc
export CGO_CFLAGS=-I/usr/local/include
export CGO_LDFLAGS=-Wl -I${ZMQ}/lib -I${GCC}/lib -lsodium -lzmq -lpthread -lstdc++ -lm -lc
export PKG_CONFIG_PATH=${ZMQ}/lib/pkgconfig

DIR_SRC=./worker
DIR_BIN=./dist
NODE_BIN=./node_modules/.bin
GO_ROOT=$(go env GOROOT)

default: build

build: install compile test

install:
	go mod vendor && go mod tidy

compile:
	go build -v -o ${DIR_BIN}/worker ${DIR_SRC}/main.go

test:
	${NODE_BIN}/mocha

fmt:
	go fmt ./ferret/... && \
	${NODE_BIN}/pretty-quick

publish:
	npm publish --access=public