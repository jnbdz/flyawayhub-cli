SHELL := bash
AppName = flyawayhub
BuildName = ${AppName}-cli
AppDir = ~/.local/bin/
AppPath = ${AppDir}${AppName}

.PHONY: all
all: build

build:
	go build -v

.PHONY: install
install:
	mv ${BuildName} ${AppPath}

.PHONY: clean
clean:
	rm -f ${BuildName}

.PHONY: uninstall
uninstall:
	rm -f ${AppPath}

.PHONY: test
test:
	go test -v ./...

.PHONY: test_coverage
test_coverage:
	go test ./... -coverprofile=coverage.out

.PHONY: dep
dep:
	go mod download

.PHONY: lint
lint:
	golangci-lint run --enable-all
