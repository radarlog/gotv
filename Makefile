.SILENT:
.DEFAULT_GOAL := linux

VERSION = $(shell git describe --tags --always --dirty --match=v*)
COMMIT = $(shell git rev-parse --short HEAD)
BUILD_TIME = $(shell date +%FT%T%z)
LDFLAGS = '-X main.VERSION=${VERSION} -X main.commit=${COMMIT} -X main.buildTime=${BUILD_TIME}'

BINARY = gotv
BUILD_DIR = $(CURDIR)/build/
CONFIG = $(CURDIR)/config.yml

M = $(shell printf "\033[32;1m▶\033[0m")

.PHONY: config
config: ; $(info $(M) copying config file)
	mkdir -p ${BUILD_DIR}
	cp -f ${CONFIG} ${BUILD_DIR}/

.PHONY: linux
linux: config ; $(info $(M) building linux binary)
	GOOS=linux GOARCH=amd64 go build -ldflags ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-linux *.go

.PHONY: raspberry
raspberry: config ; $(info $(M) building raspberry binary)
	GOOS=linux GOARCH=arm GOARM=7 go build -ldflags ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-raspberry *.go

.PHONY: clean
clean: ; $(info $(M) cleaning)
	rm -rf ${BUILD_DIR}

.PHONY: version
version:
	echo $(VERSION)
