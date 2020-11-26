# docktor

[![Build Status](https://travis-ci.com/nishgowda/docktor.svg?branch=master)](https://travis-ci.com/nishgowda/docktor)

docktor is a security tool for docker containers that allow you to implement efficient and reliable healthchecks to them.

Some highlights of docktor:
 - Creates automatic healthchecks to running docker containers
 - Implements automated healing for any unhealthy docker containers running
 - Adds manual healing as well
 - Suggests security to improvements in a given Dockerfile
 - Scans Docker images for vulnerabilites and generate reports to files

## Installation
```
go get github.com/nishgowda/docktor
```

## Building docktor
After installation building the docktor binary can be done with the following command (this will create the binary in the existing bin folder):
``` 
make build
```

However you can also install binary from this git repo:
```
go get github.com/nishgowda/docktor/bin/docktor
```
## Running docktor
 Commands
 - healthcheck
 - heal
 - autoheal
 - suggest
 - scan
 Flags
  - specify container ids --c
  - specify docker files --f
  - specify docker images --i
  - specify file to write report to --f


## Contributing
Any contributions are welcome, just clone this repo and submit a PR to this branch.