package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Valeron93/slideshower/frontend"
)

func main() {

	http.Handle("/", frontend.Handler)

	http.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	const addr = ":3000"
	slog.Info("listening", "addr", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		slog.Error("listen", "err", err)
	}
}
