package generator

import (
	"github.com/juliushill42/product-cloning-without-forks/internal/blueprint"
	"os"
	"path/filepath"
	"testing"
)

func TestGenerate(t *testing.T) {
	d := t.TempDir()
	p := blueprint.Product{Name: "X", Slug: "x-app", Pages: []blueprint.Page{{Path: "/", Title: "Home"}}, Endpoints: []blueprint.Endpoint{{Method: "GET", Path: "/api/health", Status: 200, Response: map[string]any{"ok": true}}}}
	if e := Generate(p, d); e != nil {
		t.Fatal(e)
	}
	for _, f := range []string{"go.mod", "main.go", "product.json", "README.md", "WATERMARK"} {
		if _, e := os.Stat(filepath.Join(d, f)); e != nil {
			t.Fatal(e)
		}
	}
}
