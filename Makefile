export GOPATH
export GO111MODULE=on
export QT_VERSION=5.13.2
export GOQT=$(go env GOPATH)/bin

CURRENT_OS=$(shell uname -s | awk '{print tolower($0)}')

default: build

start:
	${GOQT}/qtdeploy -fast test desktop

build:
	${GOQT}/qtdeploy build desktop

install:
	go mod vendor && go mod tidy && \
	git clone https://github.com/therecipe/env_${CURRENT_OS}_amd64_513.git vendor/github.com/therecipe/env_${CURRENT_OS}_amd64_513

generate:
	${GOQT}/qtmoc desktop

setup: install generate
	${GOQT}/qtsetup -test=false
