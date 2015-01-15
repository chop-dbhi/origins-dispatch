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

## Configuration

The `help` command will print the usage and list of subcommands. The server can be configured using command line flags or throught environment variables.

### Environment variables

- `ORIGINS_DISPATCH_DEBUG` - Set this value to `1` to turn on debugging.
- `ORIGINS_DISPATCH_SERVE_HOST` - The hostname the server will listen to.
- `ORIGINS_DISPATCH_SERVE_PORT` - The port the server will listen on.
- `ORIGINS_DISPATCH_SERVE_NEO4J` - The URL of the Neo4j endpoint the service will interact with.
