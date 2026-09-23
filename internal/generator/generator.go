package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/juliushill42/product-cloning-without-forks/internal/blueprint"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

func Generate(p blueprint.Product, out string) error {
	if err := blueprint.Validate(p); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(out, "web"), 0755); err != nil {
		return err
	}
	b, _ := json.MarshalIndent(p, "", "  ")
	if err := os.WriteFile(filepath.Join(out, "product.json"), append(b, '\n'), 0644); err != nil {
		return err
	}
	files := map[string]string{"go.mod": goMod(p), "main.go": mainGo(p), "web/index.html": indexHTML(p), "README.md": readme(p), "WATERMARK": "\":\"\n"}
	for n, c := range files {
		if err := os.WriteFile(filepath.Join(out, n), []byte(c), 0644); err != nil {
			return err
		}
	}
	return nil
}
func goMod(p blueprint.Product) string { return "module generated.local/" + p.Slug + "\n\ngo 1.22\n" }
func mainGo(p blueprint.Product) string {
	data, _ := json.Marshal(p)
	return fmt.Sprintf(`package main
import("encoding/json";"fmt";"net/http";"os";"strings";"time")
var productJSON=[]byte(%q)
type Product struct{Name string `+"`json:\"name\"`"+`;Slug string `+"`json:\"slug\"`"+`;Description string `+"`json:\"description\"`"+`;Pages []Page `+"`json:\"pages\"`"+`;Endpoints []Endpoint `+"`json:\"endpoints\"`"+`;State map[string]any `+"`json:\"state\"`"+`}
type Page struct{Path string `+"`json:\"path\"`"+`;Title string `+"`json:\"title\"`"+`;Blocks []Block `+"`json:\"blocks\"`"+`}
type Block struct{Type string `+"`json:\"type\"`"+`;Text string `+"`json:\"text,omitempty\"`"+`;Action string `+"`json:\"action,omitempty\"`"+`;Target string `+"`json:\"target,omitempty\"`"+`}
type Endpoint struct{Method string `+"`json:\"method\"`"+`;Path string `+"`json:\"path\"`"+`;Status int `+"`json:\"status\"`"+`;Response any `+"`json:\"response\"`"+`}
func main(){var p Product;if err:=json.Unmarshal(productJSON,&p);err!=nil{panic(err)};mux:=http.NewServeMux();mux.HandleFunc("GET /",func(w http.ResponseWriter,r *http.Request){if r.URL.Path!="/"{for _,pg:=range p.Pages{if pg.Path==r.URL.Path{renderPage(w,pg);return}};http.NotFound(w,r);return};renderPage(w,p.Pages[0])});for _,ep:=range p.Endpoints{ep:=ep;pattern:=strings.ToUpper(ep.Method)+" "+ep.Path;mux.HandleFunc(pattern,func(w http.ResponseWriter,r *http.Request){w.Header().Set("content-type","application/json");w.WriteHeader(ep.Status);_ = json.NewEncoder(w).Encode(ep.Response)})};addr:=os.Getenv("ADDR");if addr==""{addr="127.0.0.1:8791"};fmt.Println("generated product listening on http://"+addr);s:=&http.Server{Addr:addr,Handler:mux,ReadHeaderTimeout:5*time.Second};panic(s.ListenAndServe())}
func renderPage(w http.ResponseWriter,p Page){w.Header().Set("content-type","text/html; charset=utf-8");fmt.Fprintf(w,"<!doctype html><meta name=viewport content='width=device-width'><style>body{font:16px system-ui;background:#0a0d12;color:#f2f4f8;max-width:900px;margin:auto;padding:30px}.card{padding:18px;border:1px solid #2b3445;border-radius:14px;margin:12px 0}a,button{color:inherit}</style><h1>%%s</h1>",p.Title);for _,b:=range p.Blocks{switch b.Type{case"text":fmt.Fprintf(w,"<div class=card>%%s</div>",b.Text);case"link":fmt.Fprintf(w,"<div class=card><a href='%%s'>%%s</a></div>",b.Target,b.Text);case"button":fmt.Fprintf(w,"<div class=card><button onclick=\"location.href='%%s'\">%%s</button></div>",b.Target,b.Text)}};fmt.Fprint(w,"<footer><code>\":\"</code></footer>")}
`, string(data))
}
func indexHTML(p blueprint.Product) string {
	t := template.Must(template.New("x").Parse(`<!doctype html><html><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1"><title>{{.Name}}</title><style>body{font:16px system-ui;background:#0a0d12;color:#f4f6fa;max-width:850px;margin:auto;padding:28px}.card{border:1px solid #2a3548;border-radius:14px;padding:16px;margin:12px 0}</style></head><body><h1>{{.Name}}</h1><p>{{.Description}}</p>{{range .Pages}}<div class=card><a href="{{.Path}}">{{.Title}}</a></div>{{end}}<code>":"</code></body></html>`))
	var b bytes.Buffer
	_ = t.Execute(&b, p)
	return b.String() + "\n"
}
func readme(p blueprint.Product) string {
	return "# " + p.Name + "\n\nGenerated from an independent behavioral blueprint. No upstream repository or source tree is required.\n\nRun:\n\n```bash\ngo run .\n```\n\nWatermark: `\":\"`\n"
}
func SafeOutput(base, slug string) string { return filepath.Join(base, strings.TrimSpace(slug)) }
