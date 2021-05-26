# J&J Redirect

Mit diesem Code kann ein Docker Container erstellt werden, mit dem man auf eine URL weiterleiten kann. You can find the corresponding image [here](https://hub.docker.com/r/jjideenschmiede/jj-redirect) in Docker Hub.

## Start container

In order for a target to be defined, the REDIRECT_TARGET variable must be defined. This then works as follows.

```console
docker run --name redirect -d --restart always -e REDIRECT_TARGET=https://jj-ideenschmiede.de -p 80:80 jjideenschmiede/jj-redirect:latest
```
