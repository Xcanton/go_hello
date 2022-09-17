package main

import (
	"sync"
	"log"
	"net/http"
	"fmt"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

	/*
		handler := func(w http.ResponseWriter, r *http.Request){
		lissajous(w)
		}
		http.HandleFunc("/", handler)
	*/
	// or
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {lissajous(w)} )
	*/
}

// more complex handler
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\t%s\t%s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header [%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {  //  减小 err 的作用域，等价于先定义，后if
		log.Print(err)
	}

	for k,v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
