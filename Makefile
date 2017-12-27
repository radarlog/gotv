VERBOSE ?= 0

BINARY = gotv
PACKAGE = github.com/radarlog/${BINARY}

VERSION = $(shell git describe --tags --always --dirty --match=v*)
COMMIT = $(shell git rev-parse --short HEAD)
BUILD_TIME = $(shell date +%FT%T%z)
LDFLAGS = '-X main.VERSION=${VERSION} -X main.commit=${COMMIT} -X main.buildTime=${BUILD_TIME}'

BINPATH = $(CURDIR)/bin
GOPATH = $(CURDIR)/.gopath~
BASEPATH = $(GOPATH)/src/$(PACKAGE)
META = $(CURDIR)/meta.yml
export GOPATH

Q = $(if $(filter 1,${VERBOSE}), , @)
M = $(shell printf "\033[32;1m▶\033[0m")

.PHONY: all
all: get linux raspberry

$(BASEPATH): ; $(info $(M) setting GOPATH…)
	$Q mkdir -p $(dir $@)
	$Q ln -sf $(CURDIR) $@

.PHONY: get
get: | $(BASEPATH) ; $(info $(M) getting dependencies…)
	$Q cd $(BASEPATH) && go get

.PHONY: meta
meta: ; $(info $(M) copying meta config…)
	$Q mkdir -p ${BINPATH}
	$Q cp -f ${META} ${BINPATH}/

.PHONY: linux
linux: get meta ; $(info $(M) building linux binary…)
	GOOS=linux GOARCH=amd64 go build -ldflags ${LDFLAGS} -o ${BINPATH}/${BINARY}-linux

.PHONY: raspberry
raspberry: get meta ; $(info $(M) building raspberry binary…)
	GOOS=linux GOARCH=arm GOARM=7 go build -ldflags ${LDFLAGS} -o ${BINPATH}/${BINARY}-raspberry

.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup everything
	$Q rm -rf ${GOPATH}
	$Q rm -rf ${BINPATH}

.PHONY: version
version:
	$Q echo $(VERSION)
