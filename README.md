# Elastic APM Go agent examples

Simple examples on how to instrument your Go code using the Elastic APM Go agent.

## Setup

* `make build` to build all the examples.
* `make start` to run _Elasticsearch_, _Kibana_ and the _Elastic APM server_ in _Docker_ containers.

Other setup options:

* `make stop` to stop all the running _Docker_ containers.
* `make destroy` to stop and remove all the running _Docker_ containers.
* `make cleanup` to remove the built binaries.

## Run the Gorilla example 

* `./gorilla-example` to run a simple service based on the _Gorilla_ toolkit.
* Try calling the service few times. For instance: Using _curl_ `curl -X GET -I "http://localhost//books/MyBook/page/132"`
* Browse to your Kibana application in `http://localhost:5601/app/apm#/services/gorilla-example/`

Then you should see something like the screenshot below.

![APM on a Gorilla example - Transactions](/docs/images/apm-gorilla-example-transactions.png "APM on a Gorilla example - Transactions")

![APM on a Gorilla example - Tracing](/docs/images/apm-gorilla-example-tracing.png "APM on a Gorilla example - Tracing")

**Note:** More examples will be implemented in the future.
