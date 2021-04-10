#!/bin/bash

# run go fmt on all non formatted go files
for i in `find . -name "*.go" -type f`; do
    go fmt $i
done