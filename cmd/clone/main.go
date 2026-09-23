package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/juliushill42/product-cloning-without-forks/internal/blueprint"
	"github.com/juliushill42/product-cloning-without-forks/internal/generator"
	"github.com/juliushill42/product-cloning-without-forks/internal/server"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("clone validate|generate|serve")
		os.Exit(2)
	}
	switch os.Args[1] {
	case "validate":
		fs := flag.NewFlagSet("validate", flag.ExitOnError)
		in := fs.String("in", "examples/board-lite.json", "blueprint")
		_ = fs.Parse(os.Args[2:])
		p := load(*in)
		must(blueprint.Validate(p))
		fmt.Println("VALID", p.Name)
	case "generate":
		fs := flag.NewFlagSet("generate", flag.ExitOnError)
		in := fs.String("in", "examples/board-lite.json", "blueprint")
		out := fs.String("out", "dist/generated", "output")
		_ = fs.Parse(os.Args[2:])
		p := load(*in)
		must(generator.Generate(p, *out))
		fmt.Println(*out)
	case "serve":
		fs := flag.NewFlagSet("serve", flag.ExitOnError)
		addr := fs.String("addr", "127.0.0.1:8792", "listen")
		_ = fs.Parse(os.Args[2:])
		fmt.Println("cloner listening on http://" + *addr)
		must(http.ListenAndServe(*addr, server.Handler()))
	default:
		os.Exit(2)
	}
}
func load(p string) blueprint.Product {
	b, e := os.ReadFile(p)
	must(e)
	var x blueprint.Product
	must(json.Unmarshal(b, &x))
	return x
}
func must(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", e)
		os.Exit(1)
	}
}
