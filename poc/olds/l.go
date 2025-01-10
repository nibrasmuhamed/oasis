package main

func subUnsortz(A []int) []int {
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			return []int{-1}
		}
	}
	return nil
}

var count = 0

func Func(sr, sc, dr, dc int) {

	if sr > dr || sc > dc {
		return
	}

	if sr == dr && sc == dc {
		count++
		return
	}

	Func(sr, sc+1, dr, dc)
	Func(sr+1, sc, dr, dc)
}
