# Example Go functions

The functions are example and basic "smoke tests" for the [golang-http templates](https://github.com/openfaas/golang-http-template)

```
faas-cli template pull https://github.com/LucasRoesler/golang-http-template#feat-upgrade-to-1.18-and-use-workspaces
faas-cli build
```

## Examples:

1. [`publisher`](./publisher/) is an example `go-http` function that publishes messages to a NATS subject. This function demonstrates using
   a. external public packages: `github.com/nats-io/nats.go` and `github.com/openfaas/templates-sdk/go-http`
   b. it contains a named subpackage [`pkg/version`](./publisher/pkg/version/)
