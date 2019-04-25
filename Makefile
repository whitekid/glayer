# Makefile for nucleo-golang-api
TARGET=bin/glayer
SRC=$(shell  \
	fd -t f -e go -E '*_test.go' 2> /dev/null || \
	find . -type f -name '*.go' -not -path "./vendor/*" -not -path "*_test.go")
GO?=go

LDFLAGS = -s -w
BUILD_FLAGS?=-v -ldflags="${LDFLAGS}"

.PHONY: clean test dep tidy

all: build
build: $(TARGET)

$(TARGET): $(SRC)
	GOBIN=${PWD}/bin ${GO} install ${BUILD_FLAGS} ./...

clean:
	@rm -f ${TARGET}

test:
	LOG_LEVEL=debug ${GO} test ./...

dep:
	rm -f go.mod go.sum
	${GO} mod init github.com/whitekid/glayer
	@$(MAKE) tidy

tidy:
	${GO} mod tidy -v
