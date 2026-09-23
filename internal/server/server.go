package server

import (
	"encoding/json"
	"github.com/juliushill42/product-cloning-without-forks/internal/blueprint"
	"github.com/juliushill42/product-cloning-without-forks/internal/ui"
	"net/http"
	"time"
)

func Handler() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html; charset=utf-8")
		_, _ = w.Write(ui.Index)
	})
	m.HandleFunc("POST /api/validate", func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 2<<20)
		var p blueprint.Product
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if err := blueprint.Validate(p); err != nil {
			http.Error(w, err.Error(), 422)
			return
		}
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"valid": true, "pages": len(p.Pages), "endpoints": len(p.Endpoints), "watermark": "\":\""})
	})
	return http.TimeoutHandler(m, 10*time.Second, "timeout")
}
