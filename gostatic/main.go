/*
* Reference:
* https://codeburst.io/serving-static-files-using-a-go-web-server-d5025157a84e
 */

package main

import (
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.ListenAndServe(":8090", nil)
}
