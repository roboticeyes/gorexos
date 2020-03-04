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
VERSION := 0.1.0
BUILD := `git rev-parse HEAD`

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

INSTALL_PATH=/usr/local/bin
MAN_PATH=/usr/local/man

PKGS := $(shell go list ./... | grep -v /vendor)

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

RXC_SRC = cmd/rxc/*.go

TARGETS_LINUX = rxc
TARGETS_WINDOWS = rxc.exe
TARGETS = $(TARGETS_LINUX) $(TARGETS_WINDOWS)

all: rxc

windows: rxc-win

rxc: $(RXC_SRC)
	$(GOBUILD) -o $@ $(LDFLAGS) $(RXC_SRC)

rxc-win: $(RXC_SRC)
	GOOS=windows $(GOBUILD) -o rxc.exe $(LDFLAGS) $(RXC_SRC)

clean:
	@rm -f $(TARGETS)
	@rm -f package-v*.zip

test:
	$(GOTEST) $(PKGS)

install: all
	sudo cp -f $(TARGETS) ${INSTALL_PATH}

package: all windows
	zip package-v$(VERSION)-linux.zip $(TARGETS_LINUX) LICENSE README.md
	zip package-v$(VERSION)-win.zip $(TARGETS_WINDOWS) LICENSE README.md

uninstall:
	sudo rm -f ${INSTALL_PATH}/rxc

.PHONY: all test install uninstall
