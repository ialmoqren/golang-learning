package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Hey there</h1>
<p>It's me</p>
<p>...playing with go!</p>
<p>My name is %s but you can call me %s</p>
`, "Ibrahim Almuqrin", "Ibra")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe("", nil)
}
