package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	wg.Add(1)
	wg.Wait()
}
