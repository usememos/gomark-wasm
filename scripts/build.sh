#!/bin/sh

# exit when any command fails
set -e

cd "$(dirname "$0")/../"

tinygo build -o ./example/gomark.wasm -target wasm . 
