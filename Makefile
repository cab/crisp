NAME=crisp
PACKAGE=github.com/cab/$(NAME)
EXE=bin/$(NAME)


CWD=$(shell pwd)
PROJECT_SRC=$(shell find . -path ./vendor -prune -o -type f -name '*.go' -print)
ALL_SRC=$(shell find . -type f -name '*.go')

.PHONY: all clean run deps fmt

all: $(EXE)

clean:
	rm -rf $(EXE)

run: $(EXE)
	./$(EXE)

test: $(ALL_SRC) deps testdeps
	go test

testdeps: deps
	go get -t

deps:
	go get
	git submodule update --init

$(EXE): $(ALL_SRC) deps
	go build -o $(EXE) driver/main.go

fmt: $(PROJECT_SRC)
	@gofmt -s -l -w $(PROJECT_SRC)

watch:
	while sleep 1; do	find ./ -name '*.go' | entr -dc make test; done