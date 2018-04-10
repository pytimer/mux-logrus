mux-logrus
==============

[![GoDoc](https://godoc.org/github.com/pytimer/mux-logrus?status.svg)](https://godoc.org/github.com/pytimer/mux-logrus)

logrus middleware for groilla/mux

## Getting Started

```go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pytimer/mux-logrus"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", index).Methods(http.MethodGet)

    // add logger middleware
	r.Use(muxlogrus.NewLogger().Middleware)

	address := ":8990"
	log.Printf("Listen on address %s", address)
	http.ListenAndServe(address, r)
}
```

output log:

```sh
time="2018-04-10T11:10:39+08:00" level=info msg="completed handling request" remoteAddr=127.0.0.1 status=400 took="6.838µs"
```