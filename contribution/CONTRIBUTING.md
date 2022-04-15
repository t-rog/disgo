# Contributing

## License

You agree to license any contribution to this library under the [Apache License 2.0](#license).

## Pull Requests

Pull requests must follow the [code specification](#code-specification) and work with all [test cases](#test).

## Domain

The domain of Disgo in providing an API for HTTP/WebSocket requests. The program uses provided structures (from the [Discord API](https://discord.com/developers/docs/reference)) to provide simple abstractions for end users _(developers)_.

## Project Structure

The repository consists of a detailed [README](/README.md), [examples](/examples/), and [**API Wrapper**](/disgo/wrapper).

### Disgo

The API Wrapper consists of multiple packages. A [bundler](https://pkg.go.dev/golang.org/x/tools/cmd/bundle) is used to package the API into a single-usable package _(`disgo.go`)_.

| Package   | Description                                   |
| :-------- | :-------------------------------------------- |
| wrapper   | Contains the wrapper bundling functionality.  |
| client    | The Disgo Client Abstraction.                 |
| events    | Discord API Events.                           |
| pkg       | Utility functionality for specific libraries. |
| requests  | Discord API Requests.                         |
| resources | Discord API Resources.                        |
| sessions  | Discord API WebSocket Sessions (Gateways).    |

#### Structs

Structs are sourced from [Dasgo](https://github.com/switchupcb/dasgo) and refactored into the correct name scheme for end users. The abstraction _(i.e Resource)_ is pre-pended to the resource name _(i.e User)_ to speed up development. Modern IDE's will show the developer a list of resources when `disgo.R` is typed; rather than a bunch of irrelevant resources, functions, and variables.

```go
disgo.ResourceUser
disgo.RequestGetUser
```

#### Requests

Resource GET, DELETE, POST, PUT, BULK _(GET, ...)_ `Send()` functions are generated from the respective requests object. For more information, read [requests](../wrapper/requests/README.md).

```go
// TODO: Example of REST request function.
```

### Code Specification

#### Comments

Comments follow [Effective Go](https://golang.org/doc/effective_go#commentary) and explain why more than what _(unless the "what" isn't intuitive)_.

#### Static Code Analysis

Disgo uses [golangci-lint](https://github.com/golangci/golangci-lint) in order to statically analyze code. You can install golangci-lint with `go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.45.0` and run it using `golangci-lint run`. If you receive a `diff` error, you must add a `diff` tool in your PATH. There is one located in the `Git` bin.

If you receive `File is not ... with -...`, use `golangci-lint run --disable-all --no-config -Egofmt --fix`.

#### Fieldalignment

Disgo uses [fieldalignment]() to save memory. You can install fieldalignment with `go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest` and run it using `fieldalignment -fix <package>`.

### Test

#### Unit Tests

Unit tests are used to test logic.

#### Integration Tests

Integration tests are used to ensure functionality between the API and Discord.

# Roadmap

Disgo is currently a PROOF OF CONCEPT. Here are the steps required in order to complete it:

1. Generate **Request** code using https://github.com/switchupcb/disgo/pull/3.
2. Generate any other code.
3. Implement Client once.
4. Implement Gateway Events.
5. Implement Sharding.
6. Implement Cache (which is likely where most effort lies; caching is difficult).

In addition, we must make [decisions](/contribution/libraries/) for the following:
1. UDP connections (Voice)
2. [Audio Processing using Opus](https://discord.com/developers/docs/topics/voice-connections#encrypting-and-sending-voice)