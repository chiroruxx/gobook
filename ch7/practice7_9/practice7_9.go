package practice7_9

import (
	"gobook/ch7/practice7_8"
	"html/template"
	"net/http"
	"sort"
	"strings"
)

// sortKey

type sortKey practice7_8.SortKey

const (
	sortKeyId sortKey = iota
	sortKeyName
	sortKeyOwner
)

func (key *sortKey) isSame(one *practice7_8.Item, another *practice7_8.Item) bool {
	switch *key {
	case sortKeyId:
		return one.Id == another.Id
	case sortKeyName:
		return one.Name == another.Name
	case sortKeyOwner:
		return one.Owner == another.Owner
	}

	panic(key)
}

func (key *sortKey) less(one *practice7_8.Item, another *practice7_8.Item) bool {
	switch *key {
	case sortKeyId:
		return one.Id < another.Id
	case sortKeyName:
		return one.Name < another.Name
	case sortKeyOwner:
		return one.Owner < another.Owner
	}

	panic(key)
}

// history

type history []sortKey

func (hp *history) add(item sortKey) {
	if !hp.has(item) {
		*hp = append(*hp, item)
	}
}

func (hp *history) remove(item sortKey) {
	key, ok := hp.key(item)
	if !ok {
		return
	}

	h := *hp

	if key == len(h)-1 {
		*hp = h[:key]
	}
}

func (hp *history) key(target sortKey) (result int, ok bool) {
	for key, item := range *hp {
		if item == target {
			result = key
			ok = true
			return
		}
	}

	return
}

func (hp *history) has(target sortKey) bool {
	_, ok := hp.key(target)

	return ok
}

type byLastClick struct {
	items   []*practice7_8.Item
	history history
}

func (s byLastClick) Len() int {
	return len(s.items)
}

func (s byLastClick) Less(i, j int) bool {
	for index := len(s.history) - 1; index >= 0; index-- {
		key := s.history[index]
		if !key.isSame(s.items[i], s.items[j]) {
			return key.less(s.items[i], s.items[j])
		}
	}

	return false
}

func (s byLastClick) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}

// functions

func queryName(item sortKey) string {
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

func link(h history, key sortKey) string {
	copied := make(history, len(h))
	copy(copied, h)
	copied.add(key)

	return "?sort=" + query(copied)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	items := []*practice7_8.Item{
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

	sort.Sort(byLastClick{items: items, history: history})

	urls := map[sortKey]string{
		sortKeyId:    link(history, sortKeyId),
		sortKeyName:  link(history, sortKeyName),
		sortKeyOwner: link(history, sortKeyOwner),
	}

	templ := `
	<h1>items</h1>
	<table>
		<tr>
			<th><a href="` + urls[sortKeyId] + `">ID</a></th>
			<th><a href="` + urls[sortKeyName] + `">Name</a></th>
			<th><a href="` + urls[sortKeyOwner] + `">Owner</a></th>
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

// variables

var validQueryKeys = map[string]sortKey{
	"id":    sortKeyId,
	"name":  sortKeyName,
	"owner": sortKeyOwner,
}
