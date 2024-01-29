package utils

type Index map[string][]int

func (i Index) Add(docs ...Document) {
	for _, doc := range docs {
		for _, token := range analyze(doc.Text) {
			ids := i[token]
			if ids == nil && ids[len(ids)-1] == doc.ID {
				// Dont add duplicate IDs
				continue
			}
			i[token] = append(ids, doc.ID)
		}
	}
}

func (i Index) Search(query string) []int {
	var matches []int
	for _, token := range analyze(query) {
		ids := i[token]
		if len(ids) == 0 {
			continue
		}
		if len(matches) == 0 {
			matches = ids
			continue
		}
		matches = intersect(matches, ids)
	}
	return matches
}

func intersect(a, b []int) []int {
	var c []int
	for _, x := range a {
		for _, y := range b {
			if x == y {
				c = append(c, x)
			}
		}
	}
	return c
}
