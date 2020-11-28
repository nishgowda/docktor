# docktor

[![Build Status](https://travis-ci.com/nishgowda/docktor.svg?branch=master)](https://travis-ci.com/nishgowda/docktor)
[![Go Report Card](https://goreportcard.com/badge/github.com/nishgowda/docktor)](https://goreportcard.com/report/github.com/nishgowda/docktor)
[![codecov](https://codecov.io/gh/nishgowda/docktor/branch/master/graph/badge.svg?token=G4GHPU44W9)](https://codecov.io/gh/nishgowda/docktor)

docktor is a security tool for docker containers that allow you to implement efficient and reliable healthchecks to them.

Some highlights of docktor:
 - Creates automatic healthchecks to running docker containers
 - Implements automated healing for any unhealthy docker containers running
 - Adds manual healing as well
 - Suggests security to improvements in a given Dockerfile
 - Scans Docker images for vulnerabilites and generate reports to files

## Installation
If you would like to run the project locally use the following command
```
go get github.com/nishgowda/docktor
```

## Building docktor
After installation building the docktor binary can be done with the following command (this will create the binary in the existing bin folder)
``` 
make build
```

However you can also install binary from this git repo
```
go get github.com/nishgowda/docktor/bin/docktor
```
## Running docktor
Currently there exists 5 commands that execute each of docktor's features. 

| Command | Flags |
|---------| ------|
| healthcheck | --c (specify container id)|
|heal | --c (specify contianer id) |
|autoheal | --c (specify contianer id) |
|scan | --i (docker image) , --f (file to write report to)|
|suggest | --f (docker file) | 

**Note:** You can also use the --help command for any clarifications about the commands

## Contributing
Any contributions are welcome, just clone this repo and submit a PR to this branch.