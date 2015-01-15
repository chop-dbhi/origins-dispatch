FROM golang:onbuild

MAINTAINER Byron Ruth <b@devel.io>

ENTRYPOINT ["go-wrapper", "run"]

CMD ["--addr=0.0.0.0:5002"]

EXPOSE 5002
