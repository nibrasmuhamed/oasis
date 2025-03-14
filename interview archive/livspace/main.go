package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

const (
	chunkSize = 10 * 1024 * 1024 // 10MB
)

func main() {
	var a any
	a = 10
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
}

func lastround() {
	files := []string{
		"sorted_chunk_1.txt",
		"sorted_chunk_2.txt",
		"sorted_chunk_3.txt",
	}

	outputFile := "merged_output.txt"
	err := mergeSortedFilesWithBuffer(files, outputFile)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Merging complete. Sorted file:", outputFile)
	}
}

func mergeSortedFilesWithBuffer(files []string, outputFile string) error {
	// Open output file
	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	pq := &PriorityQueue{}
	heap.Init(pq)

	readers := make([]*bufio.Scanner, len(files))
	for i, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		reader := bufio.NewScanner(f)
		buffer := make([]byte, chunkSize)
		reader.Buffer(buffer, chunkSize)
		if reader.Scan() {
			num, _ := strconv.Atoi(reader.Text())
			heap.Push(pq, &Item{value: num, fileIndex: i})
		}
		readers[i] = reader
	}

	writer := bufio.NewWriter(output)
	defer writer.Flush()

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		fmt.Fprintln(writer, item.value)

		if readers[item.fileIndex].Scan() {
			num, _ := strconv.Atoi(readers[item.fileIndex].Text())
			heap.Push(pq, &Item{value: num, fileIndex: item.fileIndex})
		}
	}

	return nil
}

type Item struct {
	value     int
	fileIndex int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].value < pq[j].value }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
