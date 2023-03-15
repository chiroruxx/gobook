package practice7_12

import (
	"fmt"
	"gobook/ch7/http1"
	"html/template"
	"net/http"
)

type Database map[string]http1.Dollars

func (db Database) List(w http.ResponseWriter, _ *http.Request) {
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

func (db Database) Price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
