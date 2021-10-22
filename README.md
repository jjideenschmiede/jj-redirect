# J&J Redirect

[![Go](https://github.com/jjideenschmiede/jj-redirect/actions/workflows/go.yml/badge.svg)](https://github.com/jjideenschmiede/jj-redirect/actions/workflows/go.yml) [![Docker Image CI](https://github.com/jjideenschmiede/jj-redirect/actions/workflows/docker-image.yml/badge.svg)](https://github.com/jjideenschmiede/jj-redirect/actions/workflows/docker-image.yml)

This code can be used to create a Docker container that can be used to redirect to a URL. You can find the corresponding image [here](https://hub.docker.com/r/jjdevelopment/jj-redirect) in Docker Hub.

## Start container

In order for a target to be defined, the REDIRECT_TARGET variable must be defined. This then works as follows.

```console
docker run --name redirect -d --restart always -e REDIRECT_TARGET=https://jj-ideenschmiede.de -p 80:80 jjdevelopment/jj-redirect:latest
```
