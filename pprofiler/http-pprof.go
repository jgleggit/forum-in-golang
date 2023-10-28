package main

import (
	"fmt"
	"os"
	"net/http"
	_ "net/http/pprof"
	"runtime/pprof"
)

func handler(w http.ResponseWriter, r 
	*http.Request) {
	fmt.Fprintf(w, "Hello, you have requested: %s\n", r.URL.Path)
}

func main() {

	f, _ := os.Create("cpu_profile.prof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
/*
	f, _ := os.Create("mem_profile.prof")
	runtime.GC() // get up-to-date statistics
	pprof.WriteHeapProfile(f)
	f.Close()
*/
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

