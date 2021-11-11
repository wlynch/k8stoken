# k8stoken

Generate a Kubernetes client token for use anywhere!

WARNING: This is pretty hacky. This was just a fun idea I had to see if it could
work. I would recommend using client-go as much as possible for k8s client code.

This proof of concept came from wanting to use the
[client-go](https://github.com/kubernetes/client-go) libraries to get valid k8s
user credentials for use with other non-rest clients (i.e. gRPC).

Currently, the client-go implementation relies heavily on `http.RoundTripper` to
inject the user credential into requests, but this makes it difficult to only
get the token to use elsewhere. Instead of trying to refactor client-go, this
library simply constructs a fake `http.RoundTripper` and intercepts the bearer
token. This means that you get all the nice token caching behavior out of the
box!

Note: This only supports bearer tokens - cert based auth is not supported.
