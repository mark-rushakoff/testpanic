package main

import (
	"net/http"

	"github.com/elazarl/goproxy"
	"github.com/sirupsen/logrus"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	logrus.Fatal(http.ListenAndServe(":8080", proxy))
}
