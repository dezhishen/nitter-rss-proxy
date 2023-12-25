// Copyright 2021 Daniel Erat.
// All rights reserved.

package main

import (
	"github.com/derat/nitter-rss-proxy/pkg/config"
	"github.com/derat/nitter-rss-proxy/pkg/proxy"
)

func main() {
	err := config.Init("./cofig", "config")
	if err != nil {
		panic(err)
	}
	var opts proxy.Options
	err = config.Unmarshal("proxy", &opts)
	if err != nil {
		panic(err)
	}
	err = proxy.Start(opts)
	if err != nil {
		panic(err)
	}
}
