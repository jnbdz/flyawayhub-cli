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
	local goLib=~/.local/share/go/v${version}

	if [ ! -d "${goLib}" ]
	then
		echo " > Create the directory for storing Go libraries"
		mkdir -p ${goLib}
	fi

	podman run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp -v ~/.local/share/go/v${version}:/go golang:${version} go ${@}
}
execGo ${@}
