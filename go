#!/bin/bash

goVersion() {
	local localPath=$PWD
	local goModData=$(cat ${localPath}/go.mod)
	local firstRequireLine=$(echo "${goModData}" | awk '/require\s+?\(/{ print NR; exit }')
	local lastRequireLine=$(echo "${goModData}" | awk '/\)/{ print NR; exit }')
	echo "${goModData}" | sed -e "${firstRequireLine},${lastRequireLine}d" | grep -v require | grep -E "^go ([1-9])([1-9]+)?\.[1-9]([1-9]+)?$" | sed 's/go //'
}

execGo() {
	local version=$(goVersion)
	podman run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp -v ~/.local/share/go/v${version}:/go golang:${version}
}
execGo
