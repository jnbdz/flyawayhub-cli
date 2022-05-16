SHELL := bash

.PHONY: all
all: build

build:
	go build -v

.PHONY: install
install:
	mv flyawayhub-cli ~/.local/bin/flyawayhub

.PHONY: clean
clean:
	rm -f flyawayhub-cli

.PHONY: uninstall
uninstall:
	rm -f ~/.local/bin/flyawayhub
