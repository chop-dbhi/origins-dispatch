FROM golang:onbuild

MAINTAINER Byron Ruth <b@devel.io>

ENTRYPOINT ["go-wrapper", "run"]

CMD ["serve", "--host=0.0.0.0"]

EXPOSE 5002
