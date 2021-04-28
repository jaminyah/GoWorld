/*
* Reference:
* https://codeburst.io/serving-static-files-using-a-go-web-server-d5025157a84e
 */

package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Go, World!</h1>")
	fmt.Fprintf(w, "<p>Go is a wonderful language")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About page</h2>")
	fmt.Fprintf(w, "<p>Go is an opensource language</p>")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.ListenAndServe(":8090", nil)
}
