package practice7_11

import (
	"fmt"
	"gobook/ch7/http1"
	"gobook/ch7/http3"
	"net/http"
	"strconv"
	"strings"
)

type Database struct {
	http3.Database
}

func NewDatabase(items map[string]http1.Dollars) *Database {
	base := http3.Database(items)
	db := Database{Database: base}

	return &db
}

func (dbp *Database) has(item string) bool {
	_, ok := (*dbp).Database[item]
	return ok
}

func (dbp *Database) stringItem(item string) string {
	price, ok := (*dbp).Database[item]

	if !ok {
		panic("item is not found.")
	}

	return fmt.Sprintf("%s: %s\n", item, price)
}

func (dbp *Database) ListOrStore(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		dbp.store(w, req)
	} else {
		dbp.List(w, req)
	}
}

func (dbp *Database) store(w http.ResponseWriter, req *http.Request) {
	validator := validator{option: optionItemHas}
	validationErrors := validator.validate(req, dbp)

	if len(validationErrors) != 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprint(w, strings.Join(validationErrors, "\n"))
		return
	}

	item := req.FormValue("item")
	price, _ := strconv.ParseFloat(req.FormValue("Price"), 32)

	(*dbp).Database[item] = http1.Dollars(price)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, dbp.stringItem(item))
}

func (dbp *Database) OperateItem(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodDelete:
		dbp.destroy(w, req)
	case http.MethodPut:
		dbp.update(w, req)
	default:
		dbp.show(w, req)
	}
}

func (dbp *Database) show(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if !dbp.has(item) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	fmt.Fprint(w, dbp.stringItem(item))
}

func (dbp *Database) update(w http.ResponseWriter, req *http.Request) {
	validator := validator{option: optionItemNotHas}
	validationErrors := validator.validate(req, dbp)

	if len(validationErrors) != 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprint(w, strings.Join(validationErrors, "\n"))
		return
	}

	item := req.FormValue("item")
	price, _ := strconv.ParseFloat(req.FormValue("Price"), 32)

	(*dbp).Database[item] = http1.Dollars(price)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, dbp.stringItem(item))
}

func (dbp *Database) destroy(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")

	if !dbp.has(item) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	delete((*dbp).Database, item)

	w.WriteHeader(http.StatusNoContent)
}

type validationOption int

const (
	optionItemHas validationOption = iota
	optionItemNotHas
)

type validator struct {
	option validationOption
}

func (v validator) validate(req *http.Request, dbp *Database) (errors []string) {
	item := req.FormValue("item")
	price := req.FormValue("Price")

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
		errors = append(errors, "Price is required.")
	} else {
		if _, err := strconv.ParseFloat(price, 32); err != nil {
			errors = append(errors, "Price should be a number.")
		}
	}

	return
}
