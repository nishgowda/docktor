#!/bin/bash
for i in `find . -name "*.go" -type f`; do
    go fmt $i
done