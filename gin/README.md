# Gin example

Example on how to instrument a web service based on the [_Gin_](https://github.com/gin-gonic/gin) web framework using the [_Elastic APM Go agent_](https://github.com/elastic/apm-agent-go).

Please refer to the main [setup instructions](/README.md) before running the example.

## Run the example 

* `./gin-example` to run a simple service based on the _Gin_ web framework.
* Try calling the service few times. For instance: Using _curl_ `curl -X GET -I "http://localhost:8080/books/MyBook/page/132"`
* Browse to your _Kibana_ application in `http://localhost:5601/app/apm#/services/gin-example/`

Then you should see something like the screenshots below.

![APM on a Gin example - Transactions](./docs/images/apm-gin-example-transactions.png "APM on a Gin example - Transactions")

![APM on a Gin example - Tracing](./docs/images/apm-gin-example-tracing.png "APM on a Gin example - Tracing")
