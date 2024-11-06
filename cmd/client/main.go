// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/url"
	"sync"
	"os"
	"os/signal"

	goclient "goapp/internal/app/client"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	var requested_connections int
	
	flag.IntVar(&requested_connections, "c", 1, "number of connections to ws")
	flag.Parse()
	log.SetFlags(0)
	log.Println("connections:", requested_connections)

	shutdown := make(chan struct{})
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	go func() {
		<-interrupt
		log.Println("interrupt")
		close(shutdown)
	}()
	
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/goapp/ws"}
	log.Printf("connecting to %s", u.String())

	var wg sync.WaitGroup
	for i := 0; i < requested_connections; i++ {
		wg.Add(1)
		new_connection := goclient.New()
		go new_connection.Connect(i, u, shutdown, &wg)
	}
	wg.Wait()
}