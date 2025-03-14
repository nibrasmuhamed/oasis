package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
	hash := map[string][]string{}
	for _, str := range strs {
		chars := strings.Split(str, "")
		sort.Strings(chars)
		sortedStr := strings.Join(chars, "")
		_, ok := hash[sortedStr]
		if !ok {
			hash[sortedStr] = append(hash[sortedStr], str)
		} else {
			hash[sortedStr] = append(hash[sortedStr], str)
		}
	}
	out := [][]string{}
	for _, v := range hash {
		out = append(out, v)
	}
	return out
}

type Alphbetval []string

func (a Alphbetval) Len() int           { return len(a) }
func (a Alphbetval) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Alphbetval) Less(i, j int) bool { return a[i] < a[j] }
