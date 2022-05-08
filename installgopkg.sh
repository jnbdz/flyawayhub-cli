#!/bin/bash

podman run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.18 go get 
