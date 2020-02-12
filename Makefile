# Makefile for go project
#
# Author: Bernhard Reitinger
#
# Targets:
# 	all: Builds the code
# 	build: Builds the code
# 	fmt: Formats the source files
# 	clean: cleans the code
# 	install: Installs the binaries
# 	test: Runs the tests
#
VERSION := 1.0.0
BUILD := `git rev-parse HEAD`

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

PKGS := $(shell go list ./... | grep -v /vendor)

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

RXC_SRC = cmd/rxc/*.go

TARGETS = rxc

all: rxc

rxc: $(RXC_SRC)
	$(GOBUILD) -o $@ $(LDFLAGS) $(RXC_SRC)

clean:
	@rm -f $(TARGETS)

test:
	$(GOTEST) $(PKGS)

install: all
	mkdir -p ${DESTDIR}${PREFIX}/bin
	cp -f $(TARGETS) ${DESTDIR}${PREFIX}/bin
	mkdir -p ${DESTDIR}${MANPREFIX}/man1
	sed "s/VERSION/${VERSION}/g" < rxc.1 > ${DESTDIR}${MANPREFIX}/man1/rxc.1
	chmod 644 ${DESTDIR}${MANPREFIX}/man1/rxc.1

uninstall:
	rm -f ${DESTDIR}${PREFIX}/bin/rxc\
		${DESTDIR}${MANPREFIX}/man1/rxc.1

.PHONY: all test install uninstall
