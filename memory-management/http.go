package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	http.HandleFunc("/bar", func (w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.Path))
	})

	http.ListenAndServe(":8080", nil)
}

/*
USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
vagrant  27765  0.0  0.1 746636  6744 pts/0    Sl+  16:55   0:00 ./http
 */