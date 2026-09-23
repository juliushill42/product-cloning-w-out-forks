package blueprint

import (
	"errors"
	"fmt"
	"regexp"
)

var slugRE = regexp.MustCompile(`^[a-z0-9][a-z0-9-]{1,62}$`)

type Product struct {
	Name        string         `json:"name"`
	Slug        string         `json:"slug"`
	Description string         `json:"description"`
	Pages       []Page         `json:"pages"`
	Endpoints   []Endpoint     `json:"endpoints"`
	State       map[string]any `json:"state"`
}
type Page struct {
	Path   string  `json:"path"`
	Title  string  `json:"title"`
	Blocks []Block `json:"blocks"`
}
type Block struct {
	Type   string `json:"type"`
	Text   string `json:"text,omitempty"`
	Action string `json:"action,omitempty"`
	Target string `json:"target,omitempty"`
}
type Endpoint struct {
	Method   string `json:"method"`
	Path     string `json:"path"`
	Status   int    `json:"status"`
	Response any    `json:"response"`
}

func Validate(p Product) error {
	if p.Name == "" || !slugRE.MatchString(p.Slug) {
		return errors.New("valid name and slug required")
	}
	seen := map[string]bool{}
	for _, pg := range p.Pages {
		if pg.Path == "" || pg.Path[0] != '/' {
			return fmt.Errorf("invalid page path %q", pg.Path)
		}
		if seen[pg.Path] {
			return fmt.Errorf("duplicate page %s", pg.Path)
		}
		seen[pg.Path] = true
	}
	for _, e := range p.Endpoints {
		if e.Path == "" || e.Path[0] != '/' {
			return fmt.Errorf("invalid endpoint %q", e.Path)
		}
		if e.Status < 100 || e.Status > 599 {
			return fmt.Errorf("invalid endpoint status %d", e.Status)
		}
	}
	return nil
}
