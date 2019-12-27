# Elastic APM Go agent examples

Simple examples on how to instrument your Go code using the Elastic APM Go agent.

## Setup

* `make build` to build all the examples.
* `make start` to run _Elasticsearch_, _Kibana_ and the _Elastic APM server_ in _Docker_ containers.

Other setup options:

* `make stop` to stop all the running _Docker_ containers.
* `make destroy` to stop and remove all the running _Docker_ containers.
* `make cleanup` to remove the built binaries.

## Examples 

* [Gorilla](/gorilla/README.md)
* [Gin](/gin/README.md)

**Note:** More examples will be implemented in the future.
