package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/Valeron93/slideshower/html"
)

func main() {

	http.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		if err := html.Templates.ExecuteTemplate(w, "index.html", nil); err != nil {
			log.Print(err)
		}
	})

	http.Handle("/static/", html.StaticHandler)

	const addr = ":3000"
	slog.Info("listening", "addr", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		slog.Error("listen", "err", err)
	}
}
