export SHELL:=/usr/bin/env bash -O extglob -c
export GO15VENDOREXPERIMENT:=1

build: GOOS ?= darwin
build: GOARCH ?= amd64
build:
	rm -f jsonify
	GOOS=${GOOS} GOARCH=${GOARCH} go build .

release-linux:
	GOOS=linux $(MAKE) build
	file jsonify
	tar Jcf jsonify-`git describe --abbrev=0 --tags`-linux-amd64.txz jsonify

release-darwin:
	GOOS=darwin $(MAKE) build
	file jsonify
	tar Jcf jsonify-`git describe --abbrev=0 --tags`-darwin-amd64.txz jsonify

release: clean release-linux release-darwin

clean:
	rm -f jsonify
	rm -f jsonify-*.txz

run: build
	./jsonify
