# ci-server
> Simple continuous integration server for automatically deploying
> [jadeocr-next](https://next.jadeocr.com)

This is a simple server that builds [jadeocr-next](https://next.jadeocr.com) in
Docker upon pushes to the jadeocr-next
[repository](https://github.com/jadeocr/jadeocr-next).

Using GitHub webhooks on the jadeocr-next repository, this server triggers
shell commands to fetch the changes from GitHub, rebuild the docker-compose
images, and start the new containers.

## Get Started
To start the development server, first clone this repository, then run `make`
or `make dev` in the root of the repo. 

## Building
To build the executable, run `make build` in the root of the repo, which will
place the binary in the root of the repo as `ci-server.sh`.

## Using
To run the CI server, copy `ci-server.sh` into the root of the
[jadeocr-next](https://github.com/jadeocr/jadeocr-next) repository and run as a
detached process.
