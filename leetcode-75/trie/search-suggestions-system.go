package trie

import (
	"sort"
	"strings"
)

func suggestedProducts(products []string, searchWord string) [][]string {
	res := make([][]string, 0)
	if products == nil {
		return res
	}
	sort.Strings(products)
	for i := 1; i <= len(searchWord); i++ {
		res = append(res, searchWords(products, searchWord[:i]))
	}
	return res
}

func searchWords(products []string, word string) []string {
	res := make([]string, 0)
	for i := 0; i < len(products); i++ {
		if strings.HasPrefix(products[i], word) {
			res = append(res, products[i])
		}
		if len(res) == 3 {
			break
		}
	}
	return res
}
