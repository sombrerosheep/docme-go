BUILDDIR=./build
ALLSRC=$(shell find . -type f -name "*.go")
APP_TARGET=app

.PHONY: clean build-docme build fmt-check fmt-apply fmt test deps

clean:
	rm $(BUILDDIR)/*

build-docme:
	go build -o $(BUILDDIR)/docme ./cmd/docme

build: build-docme

fmt-check:
	goimports -l -d -e $(ALLSRC)

fmt-apply:
	goimports -l -w $(ALLSRC)

fmt: fmt-apply

test:
	go test -v ./...

deps:
	go mod download
