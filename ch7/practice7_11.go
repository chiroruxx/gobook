package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	db := database4{"shoes": 50, "socks": 5}
	http.HandleFunc("/", db.listOrStore)
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/item", db.operateItem)
	http.HandleFunc("/price", db.price)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars4 float32

func (d dollars4) String() string { return fmt.Sprintf("$%.2f", d) }

type database4 map[string]dollars4

func (dbp *database4) has(item string) bool {
	_, ok := (*dbp)[item]
	return ok
}

func (dbp *database4) stringItem(item string) string {
	price, ok := (*dbp)[item]

	if !ok {
		panic("item is not found.")
	}

	return fmt.Sprintf("%s: %s\n", item, price)
}

func (dbp *database4) listOrStore(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		dbp.store(w, req)
	} else {
		dbp.list(w, req)
	}
}

func (dbp *database4) list(w http.ResponseWriter, _ *http.Request) {
	for item := range *dbp {
		fmt.Fprint(w, dbp.stringItem(item))
	}
}

func (dbp *database4) store(w http.ResponseWriter, req *http.Request) {
	validator := validator{option: optionItemHas}
	validationErrors := validator.validate(req, dbp)

	if len(validationErrors) != 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprint(w, strings.Join(validationErrors, "\n"))
		return
	}

	item := req.FormValue("item")
	price, _ := strconv.ParseFloat(req.FormValue("price"), 32)

	(*dbp)[item] = dollars4(price)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, dbp.stringItem(item))
}

func (dbp *database4) operateItem(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodDelete:
		dbp.destroy(w, req)
	case http.MethodPut:
		dbp.update(w, req)
	default:
		dbp.show(w, req)
	}
}

func (dbp *database4) show(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if !dbp.has(item) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	fmt.Fprint(w, dbp.stringItem(item))
}

func (dbp *database4) update(w http.ResponseWriter, req *http.Request) {
	validator := validator{option: optionItemNotHas}
	validationErrors := validator.validate(req, dbp)

	if len(validationErrors) != 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprint(w, strings.Join(validationErrors, "\n"))
		return
	}

	item := req.FormValue("item")
	price, _ := strconv.ParseFloat(req.FormValue("price"), 32)

	(*dbp)[item] = dollars4(price)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, dbp.stringItem(item))
}

func (dbp *database4) destroy(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")

	if !dbp.has(item) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	delete(*dbp, item)

	w.WriteHeader(http.StatusNoContent)
}

func (dbp *database4) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := (*dbp)[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

type validationOption int

const (
	optionItemHas validationOption = iota
	optionItemNotHas
)

type validator struct {
	option validationOption
}

func (v validator) validate(req *http.Request, dbp *database4) (errors []string) {
	item := req.FormValue("item")
	price := req.FormValue("price")

	var itemExistValidation func(item string) (bool, string)
	switch v.option {
	case optionItemHas:
		itemExistValidation = func(item string) (bool, string) {
			return dbp.has(item), "item is already exists."
		}
	case optionItemNotHas:
		itemExistValidation = func(item string) (bool, string) {
			return !dbp.has(item), "item does not exist."
		}
	default:
		itemExistValidation = func(item string) (bool, string) {
			return false, ""
		}
	}

	if result, message := itemExistValidation(item); result {
		errors = append(errors, message)
	} else if item == "" {
		errors = append(errors, "item is required.")
	}

	if price == "" {
		errors = append(errors, "price is required.")
	} else {
		if _, err := strconv.ParseFloat(price, 32); err != nil {
			errors = append(errors, "price should be a number.")
		}
	}

	return
}
