package practice5_10

func TopoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	var visit func(item string)
	visit = func(item string) {
		if !seen[item] {
			seen[item] = true
			for _, dependent := range m[item] {
				visit(dependent)
			}
			order = append(order, item)
		}
	}

	visitALl := func(m map[string][]string) {
		for item := range m {
			visit(item)
		}
	}

	visitALl(m)
	return order
}
