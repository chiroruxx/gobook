package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"
)

// Item

type Item2 struct {
	Id    uint
	Name  string
	Owner string
}

// SortKey

type SortKey2 int

const (
	SortKey2Id SortKey2 = iota
	SortKey2Name
	SortKey2Owner
)

func (key *SortKey2) isSame(one *Item2, another *Item2) bool {
	switch *key {
	case SortKey2Id:
		return one.Id == another.Id
	case SortKey2Name:
		return one.Name == another.Name
	case SortKey2Owner:
		return one.Owner == another.Owner
	}

	panic(key)
}

func (key *SortKey2) less(one *Item2, another *Item2) bool {
	switch *key {
	case SortKey2Id:
		return one.Id < another.Id
	case SortKey2Name:
		return one.Name < another.Name
	case SortKey2Owner:
		return one.Owner < another.Owner
	}

	panic(key)
}

// History

type history []SortKey2

func (hp *history) add(item SortKey2) {
	if !hp.has(item) {
		*hp = append(*hp, item)
	}
}

func (hp *history) remove(item SortKey2) {
	key, ok := hp.key(item)
	if !ok {
		return
	}

	h := *hp

	if key == len(h)-1 {
		*hp = h[:key]
	}
}

func (hp *history) key(target SortKey2) (result int, ok bool) {
	for key, item := range *hp {
		if item == target {
			result = key
			ok = true
			return
		}
	}

	return
}

func (hp *history) has(target SortKey2) bool {
	_, ok := hp.key(target)

	return ok
}

type byLastClick2 struct {
	Items   []*Item2
	history history
}

func (s byLastClick2) Len() int {
	return len(s.Items)
}

func (s byLastClick2) Less(i, j int) bool {
	for index := len(s.history) - 1; index >= 0; index-- {
		key := s.history[index]
		if !key.isSame(s.Items[i], s.Items[j]) {
			return key.less(s.Items[i], s.Items[j])
		}
	}

	return false
}

func (s byLastClick2) Swap(i, j int) {
	s.Items[i], s.Items[j] = s.Items[j], s.Items[i]
}

// functions

func queryName(item SortKey2) string {
	for key, value := range validQueryKeys {
		if value == item {
			return key
		}
	}

	panic(item)
}

func query(h history) string {
	strs := make([]string, 0)
	for _, key := range h {
		strs = append(strs, queryName(key))
	}

	return strings.Join(strs, ",")
}

func link(h history, key SortKey2) string {
	copied := make(history, len(h))
	copy(copied, h)
	copied.add(key)

	return "?sort=" + query(copied)
}

// variables

var validQueryKeys = map[string]SortKey2{
	"id":    SortKey2Id,
	"name":  SortKey2Name,
	"owner": SortKey2Owner,
}

// main

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	items := []*Item2{
		{1, "pencil", "Mark"},
		{2, "pencil", "John"},
		{3, "joint", "John"},
		{4, "joint", "Mark"},
	}

	query := r.URL.Query().Get("sort")
	keyNames := strings.Split(query, ",")
	history := make(history, 0)
	for _, keyName := range keyNames {
		key, ok := validQueryKeys[keyName]
		if ok {
			history.add(key)
		}
	}

	sort.Sort(byLastClick2{Items: items, history: history})

	urls := map[SortKey2]string{
		SortKey2Id:    link(history, SortKey2Id),
		SortKey2Name:  link(history, SortKey2Name),
		SortKey2Owner: link(history, SortKey2Owner),
	}

	templ := `
	<h1>Items</h1>
	<table>
		<tr>
			<th><a href="` + urls[SortKey2Id] + `">ID</a></th>
			<th><a href="` + urls[SortKey2Name] + `">Name</a></th>
			<th><a href="` + urls[SortKey2Owner] + `">Owner</a></th>
		</tr>
		{{range .}}
			<tr>
				<td>{{.Id}}</td>
				<td>{{.Name}}</td>
				<td>{{.Owner}}</td>
			</tr>
		{{end}}
	</table>
`
	t := template.Must(template.New("issues").Parse(templ))
	t.Execute(w, items)
}
