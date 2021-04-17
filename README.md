# docktor

[![Build Status](https://travis-ci.com/nishgowda/docktor.svg?branch=master)](https://travis-ci.com/nishgowda/docktor)
[![Go Report Card](https://goreportcard.com/badge/github.com/nishgowda/docktor)](https://goreportcard.com/report/github.com/nishgowda/docktor)
[![codecov](https://codecov.io/gh/nishgowda/docktor/branch/master/graph/badge.svg?token=G4GHPU44W9)](https://codecov.io/gh/nishgowda/docktor)

docktor is a security tool that allows you to secure and implement reliable healthchecks on running Docker containers.

Some highlights of docktor:
 - Creates automatic healthchecks for running containers
 - Implements automated healing for unhealthy running containers 
 - Adds manual healing as well
 - Suggests security improvements in a given Dockerfile
 - Scans Docker images for vulnerabilites and generate reports to files

## Building docktor
If you would like to run the project locally clone the repo. 

After installation, building the docktor binary can be done with the following command (this will create the binary in the existing bin folder)
``` 
make build
```

However you can also install binary from this git repo
```
go get github.com/nishgowda/docktor/bin/docktor
```

## Command Line
You can run the features of docktor in the command line by running the executable
### Usage
```
# run healthcheck on nginx container
./docktor healtheck --c nginx

# heal an unhealthy container named ng
./docktor heal --c ng

# apply autoheal to container ng
./doctkor autoheal --c ng

# scan for vulnerabilites in nginx container and write output to file location 
./doctkor scan --i nginx --f data/ouptut.txt

# suggest improvemets for dockerfile
./docktor suggest --f Dockerfile
```
**NOTE:** You must enable [Snyk](https://snyk.io/blog/snyk-docker-secure-containerized-applications/) to use the scan feature.

## API
You can also start the server and make requests to perform docktor functions.
 - GET /hcheck?containers=
 - GET /heal?containers= 
 - GET /aheal?containers=
- GET /scan?image=foo&file=bar
- GET /suggest?file=

### Usage
```
# start the server default port is 3001
./docktor server --p 3001

# If you dont specify the container it will perform functions on all running containers
curl --request GET 'http://localhost:3001/hcheck?containers=nginx'

curl --request GET 'http://localhost:3001/heal?containers=ngninx'

curl --request GET 'http://localhost:3001/aheal?containers=ng'

# must specify the docker image 
curl --request GET 'http://localhost:3001/scan?image=nginx'

# must specify the location of the docker file 
curl --request GET 'http://localhost:3001/suggest?file=Dockerfile
```

## Contributing
Any contributions are welcome, you're welcome to add an issue or just clone this repo and submit a PR to this branch.
