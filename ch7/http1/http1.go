package http1

import (
	"fmt"
	"net/http"
)

type Dollars float32

func (d Dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type Database map[string]Dollars

func (db Database) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
