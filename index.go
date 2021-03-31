package main

type index map[string][]int

func (idx index) feed(docs []document) {
	for _, doc := range docs {
		for _, token := range analyze(doc.Text) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}

func (idx index) search(text string) [][]int {
	var r [][]int
	for _, token := range analyze(text) {
		if ids, ok := idx[token]; ok {
			r = append(r, ids)
		}
	}
	return r
}
