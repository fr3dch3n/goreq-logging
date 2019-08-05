# goreq-logging
> Log requests in a go webserver.

[![Build Status][travis-image]][travis-url]

THis library will log all HTTP-Requests in a golang webserver. 
This comes in handy if you want global request-logging enabled.
Building the log-messages is done via [logrus](https://github.com/sirupsen/logrus).

## Installation

* with go-get: `go get github.com/fr3dch3n/goreq-logging`
* with dep: `dep ensure --add github.com/fr3dch3n/goreq-logging`


## Usage example

**Enable JSON-logging**
```go
package main

import (
	logHandler "github.com/fr3dch3n/goreq-logging"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	err := http.ListenAndServe(":80", logHandler.LogRequests(http.DefaultServeMux))
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
```

## Release History

* 0.0.1
    * initial release

## Meta

[@fr3dch3n](https://twitter.com/fr3dch3n)

Distributed under the Apache 2.0 license. See ``LICENSE`` for more information.

## Contributing

1. Fork it (<https://github.com/fr3dch3n/goreq-logging/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request

<!-- Markdown link & img dfn's -->
[travis-image]: https://img.shields.io/travis/fr3dch3n/goreq-logging/master.svg?style=flat-square
[travis-url]: https://travis-ci.org/fr3dch3n/goreq-logging
