package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	db := database5{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars5 float32

func (d dollars5) String() string { return fmt.Sprintf("$%.2f", d) }

type database5 map[string]dollars5

func (db database5) list(w http.ResponseWriter, _ *http.Request) {
	templateString := `
<!DOCTYPE html>
<html lang="en">
<body>
<dl>
{{ range $item, $price := . }}
  <dt>{{ $item }}</dt>
  <dd>{{ $price }}</dd>
{{ end }}
</body>
</html>
`

	templateEngine := template.Must(template.New("itemList").Parse(templateString))
	templateEngine.Execute(w, db)
}

func (db database5) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
