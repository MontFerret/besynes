export GOPATH
export GO111MODULE=on

CURRENT_OS=$(shell uname -s | awk '{print tolower($0)}')

default: build

start:
	qtdeploy -fast test desktop

build:
	qtdeploy build desktop

install:
	go mod vendor && go mod tidy && \
	git clone https://github.com/therecipe/env_${CURRENT_OS}_amd64_513.git vendor/github.com/therecipe/env_${CURRENT_OS}_amd64_513

generate:
	qtmoc desktop

setup: install generate
	qtesetup -test=false
