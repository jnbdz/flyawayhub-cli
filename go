#!/bin/bash

goVersion() {
	local localPath=$PWD
	local goModData=$(cat ${localPath}/go.mod)
	local firstRequireLine=$(echo "${goModData}" | awk '/require\s+?\(/{ print NR; exit }')
	echo "${goModData}" | sed -e '7,10d' | grep -v require | grep -E "^go ([1-9])([1-9]+)?\.[1-9]([1-9]+)?$" | sed 's/go //'
}

execGo() {
	local version=$(goVersion)
	echo podman run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp -v ~/.local/share/go:/go golang:${version}
}
execGo
