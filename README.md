# Example Go functions

The functions are example and basic "smoke tests" for the [golang-http templates](https://github.com/openfaas/golang-http-template)

```
faas-cli template pull https://github.com/LucasRoesler/golang-http-template#feat-automatic-vendor-detection
faas-cli build
```

## Examples:

1. [`publisher`](./publisher/) is an example `go-http` function that publishes messages to a NATS subject. This function demonstrates using
   a. external public packages: `github.com/nats-io/nats.go` and `github.com/openfaas/templates-sdk/go-http`
   b. it contains a named subpackage [`pkg/version`](./publisher/pkg/version/)

2. [`quoter`](./quoter/) is an example `go-http-middleware` function that echos the input in quotes. This function demonstrates using
   a. local replace statements in the handler `go.mod` file
   b. a function that logs

3. [`jules`](./jules/) is an example `go-http-middleware` function that echos the user input like Jules from Pulp Fiction: "Say {input} one more time!". This function demonstrates
   a. using vendoring
   b. a function that logs
