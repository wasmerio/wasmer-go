#!/bin/bash
rm -r vendor || echo 'error ignored'
mkdir -p vendor/github.com/wasmerio/wasmer-go
cp -r ../wasmer vendor/github.com/wasmerio/wasmer-go
GODEBUG=cgocheck=2 go build -o app
ldd app || echo 'error ignored'
otool -L app || echo 'error ignored'
./app
