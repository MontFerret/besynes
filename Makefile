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

DIR_WORKER_SRC=./worker
DIR_APP_SRC=./app
DIR_BIN=./dist
DIR_PROTO=./proto
NODE_BIN=./node_modules/.bin
GO_ROOT=$(go env GOROOT)

default: build

build: install compile test

install:
	go mod vendor && go mod tidy

compile:
	go build -v -o ${DIR_BIN}/worker ${DIR_WORKER_SRC}/main.go

generate:
	rm -rf ${DIR_WORKER_SRC}/app/messaging/dto/* ${DIR_APP_SRC}/main/modules/execution/dto/* && \
	protoc \
		--proto_path=${DIR_PROTO} \
		--go_out=${DIR_WORKER_SRC}/app/messaging/dto \
		--js_out=import_style=commonjs,binary:${DIR_APP_SRC}/main/modules/execution/dto \
		${DIR_PROTO}/*.proto

test:
	${NODE_BIN}/mocha

fmt:
	go fmt ./ferret/... && \
	${NODE_BIN}/pretty-quick

publish:
	npm publish --access=public