package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr, r.Method, r.RequestURI, r.UserAgent())
	defer func(t time.Time) {
		log.Println(r.RemoteAddr, time.Since(t), "wasted")
	}(time.Now())

	flusher, ok := w.(http.Flusher)
	if !ok {
		code := http.StatusTeapot
		http.Error(w, http.StatusText(code), code)
		return
	}

	w.Header().Set("X-Content-Type-Options", "nosniff")

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		select {
		case <-r.Context().Done():
			return
		default:
			w.Write([]byte("💩"))
			flusher.Flush()
		}
	}
}

func main() {
	listen := flag.String("listen", ":8080", "listen at")
	flag.Parse()

	http.HandleFunc("/", handler)
	log.Println("listen at", *listen)
	log.Fatal(http.ListenAndServe(*listen, nil))
}
