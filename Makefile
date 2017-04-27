
GOFILES_NOVENDOR = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
UNITTEST_PACKAGES = $(shell go list ./... | grep -v /vendor/)


all: fmt vet build ut

fmt:
	gofmt -l -w ${GOFILES_NOVENDOR}

vet:
	go vet ${UNITTEST_PACKAGES}

build:
	go build -ldflags -s -o bin/file_writer examples/file_writer.go

test:
	make ut

ut:
	go test -ldflags -s -v --cover ${UNITTEST_PACKAGES}

