# Origins Dispatch

[![Build Status](https://travis-ci.org/chop-dbhi/origins-dispatch.svg?branch=master)](https://travis-ci.org/chop-dbhi/origins-dispatch)

A service receives payloads from the Origins service (currently via a POST request) and dispatches the message to registered webhooks and built-in notification handlers.

## Docker

Default. Runs the service on port 5002.

```
docker run -d --name origins-dispatch -P dbhi/origins-dispatch
```

Print help and usage.

```
docker run dbhi/origins-dispatch help
```
