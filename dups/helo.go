package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	x := 5332
	as := split(x)
	ds := split(x)
	sort.Slice(as, func(i, j int) bool {
		return as[i] < as[j]
	})
	sort.Slice(ds, func(i, j int) bool {
		return ds[i] > ds[j]
	})
	i := SliceToInt(as)
	j := SliceToInt(ds)

	rec(i, j, 0)
}

func SliceToInt(a Asc) int {
	old := ""
	for _, val := range a {
		x := strconv.Itoa(val)
		newString := old + x
		old = newString
	}
	x, _ := strconv.Atoi(old)

	return x
}

func rec(as, ds, i int) {
	if i == 10 {
		return
	}
	i++
	res := ds - as
	fmt.Println(res)
	rec(ds, as, i)
}

type Asc []int

func (a Asc) Len() int           { return len(a) }
func (a Asc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Asc) Less(i, j int) bool { return a[i] < a[j] }

func split(n int) Asc {
	a := Asc{}
	for n > 0 {
		x := n % 10
		a = append(a, x)
		n = (n - x) / 10
	}
	return a

}
