package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"sync"
)

var (
	numBurn        int
	updateInterval int
)

func cpuBurn(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			for i := 0; i < 2147483647; i++ {
			}
		}
	}
}

var (
	runningMx sync.Mutex
	running   bool
	cancel    func()
)

func on(w http.ResponseWriter, r *http.Request) {

	runningMx.Lock()
	defer runningMx.Unlock()
	if running {
		fmt.Fprint(w, "already running, call /off first")
		return
	}
	running = true

	nRaw := r.URL.Query().Get("n")
	n, _ := strconv.Atoi(nRaw)
	if n <= 0 {
		n = runtime.NumCPU()
	}
	runtime.GOMAXPROCS(n)
	var ctx context.Context
	ctx, cancel = context.WithCancel(context.Background())

	for i := 0; i < n; i++ {
		go cpuBurn(ctx)
	}
	log.Printf("started with %d CPU\n", n)

	fmt.Fprintf(w, "started with %d cores usage", n)
}

func off(w http.ResponseWriter, r *http.Request) {
	runningMx.Lock()
	defer runningMx.Unlock()

	if cancel != nil {
		cancel()
	}

	running = false
	fmt.Fprint(w, "finished")
	log.Println("finished")
}

func main() {
	port := flag.String("p", "8080", "port")
	flag.Parse()

	http.HandleFunc("/on", on)
	http.HandleFunc("/off", off)
	http.ListenAndServe(":"+*port, nil)
}
