package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"sort"

	"fmt"

	"github.com/gorilla/mux"
)

// func trap(A []int) int {
// 	sum, l, c := 0, 0, 0
// 	for i := 0; i < len(A); i++ {
// 		if A[i] > 0 {
// 			l = A[i]
// 			break
// 		}
// 	}
// 	for i := l; i < len(A); i++ {

//		}
//	}
func singleNumber(A []int) int {
	numarray := make([]int, 32)
	for i := 1; i < 33; i++ {
		c := 0
		for j := 0; j < len(A); j++ {
			if (A[j]>>(i-1))&1 == 1 {
				c++
			}
		}
		if c%3 == 1 {
			numarray[i-1] = c % 3
		}
	}
	return binaryArrayToInt(numarray)
}
func binaryArrayToInt(binaryArray []int) int {
	result := 0
	// reverseSlice(binaryArray)

	for i, bit := range binaryArray {
		result += bit * int(math.Pow(float64(2), float64(i)))
	}

	return int(result)
}
func reverseSlice(s []int) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}
}

// func main() {
// 	args := os.Args[1:]

// 	// Check if there is at least one argument
// 	if len(args) < 1 {
// 		fmt.Println("Usage: go run yourfile.go <string1> <string2> ...")
// 		return
// 	}

// 	// Create a new router
// 	router := mux.NewRouter()

// 	// Define a handler function for the "/hello" route
// 	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		// Send the specified strings as the response
// 		fmt.Fprintf(w, "Message to you: %v\n", args)
// 		fmt.Fprintf(w, "you can send message using this command:")
// 		fmt.Fprintf(w, "curl -H \"message: [message_value_here]\" http://[replace_the_url]/send")

// 	})
// 	router.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
// 		// Send the specified strings as the response
// 		x := r.Header.Get("message")
// 		fmt.Println("messager: ", x)
// 	}).Methods("POST")
// 	// Start the HTTP server on port 8080
// 	http.Handle("/", router)
// 	l := log.Default()
// 	l.Println("running in 1000")
// 	http.ListenAndServe(":1000", nil)
// }

// func main() {
// 	var wg sync.WaitGroup

// 	// Create a wait group to wait for the goroutines to finish
// 	wg.Add(2)

// 	// Start a goroutine to listen for messages from a Kafka topic
// 	go func() {
// 		defer wg.Done()

// 		consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
// 			"bootstrap.servers": "localhost:9092",
// 			"group.id":          "my-group",
// 			"auto.offset.reset": "earliest",
// 		})

// 		if err != nil {
// 			fmt.Printf("Failed to create consumer: %s\n", err)
// 			return
// 		}
// 		defer consumer.Close()

// 		consumer.SubscribeTopics([]string{"usage_engine_kafka_message_topic"}, nil)

// 		for {
// 			ev := consumer.Poll(0)
// 			switch e := ev.(type) {
// 			case *kafka.Message:
// 				fmt.Printf("Received message on topic %s: %s\n", *e.TopicPartition.Topic, string(e.Value))
// 				// Process the message as needed
// 			case kafka.Error:
// 				fmt.Printf("Error: %v\n", e)
// 			}
// 		}
// 	}()

// 	// Start another goroutine to send messages to a Kafka topic
// 	go func() {
// 		defer wg.Done()

// 		producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
// 		if err != nil {
// 			fmt.Printf("Failed to create producer: %s\n", err)
// 			return
// 		}
// 		defer producer.Close()

// 		topic := "another_kafka_topic"

// 		for {
// 			var message string
// 			fmt.Print("Enter a message to send (or type 'exit' to quit): ")
// 			fmt.Scanln(&message)

// 			if message == "exit" {
// 				break
// 			}

// 			producer.Produce(&kafka.Message{
// 				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
// 				Value:          []byte(message),
// 			}, nil)
// 		}
// 	}()

// 	// Wait for both goroutines to finish
// 	wg.Wait()
// }

type NrfServer struct {
	// TODO: Implement this field to store the NRF server state.
}

func (nrf *NrfServer) RegisterNf(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this function to register a new NF instance.
	var nfInstance NfInstance
	err := json.NewDecoder(r.Body).Decode(&nfInstance)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid NF instance data: %v", err)
		return
	}

	// TODO: Validate the NF instance data.

	// TODO: Create the NF instance in the database.

	// Write a success response to the client.
	log.Println("successful registration.", http.StatusCreated, nfInstance)
	nfInstance.HeartBeatTimer = 20
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nfInstance)
}

func (nrf *NrfServer) UpdateNf(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this function to update an NF instance.
	nfInstanceId := mux.Vars(r)["nfInstanceId"]
	log.Println(nfInstanceId)
	var nfInstance NfInstance
	err := json.NewDecoder(r.Body).Decode(&nfInstance)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid NF instance data: %v", err)
		return
	}

	// TODO: Validate the NF instance data.

	// TODO: Update the NF instance in the database.

	// Write a success response to the client.
	log.Println("successful update.", http.StatusOK, nfInstance.HeartBeatTimer)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(nfInstance)
}

func (nrf *NrfServer) DeleteNf(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this function to delete an NF instance.
	nfInstanceId := mux.Vars(r)["nfInstanceId"]
	log.Println(nfInstanceId)
	// TODO: Delete the NF instance from the database.

	// Write a success response to the client.
	log.Println("successful de-registration.", http.StatusNoContent)

	w.WriteHeader(http.StatusNoContent)
}

func xo() {
	nrf := &NrfServer{}

	r := mux.NewRouter()

	// Add routes for each NRF API.
	r.HandleFunc("/nnrf-nfm/v1/nf-instances/{nfInstanceId}", nrf.RegisterNf).Methods("PUT")
	r.HandleFunc("/nnrf-nfm/v1/nf-instances/{nfInstanceId}", nrf.UpdateNf).Methods("PATCH")
	r.HandleFunc("/nnrf-nfm/v1/nf-instances/{nfInstanceId}", nrf.DeleteNf).Methods("DELETE")

	// Start the server.
	fmt.Println("Starting NRF server on port http://localhost:5050")
	http.ListenAndServe(":5050", r)
}

// Implement the remaining functions to complete the NRF server.

func mergeSoted(nums1 []int, m int, nums2 []int, n int) {
	out := []int{}
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			out = append(out, nums1[i])
			i++
		} else {
			out = append(out, nums2[j])
			j++
		}
	}
	for i < m {
		out = append(out, nums1[i])
		i++
	}
	for j < n {
		out = append(out, nums2[j])
		j++
	}
	nums1 = out
	// for i := 0 ; i<len(out); i++{
	//   if i < m {
	//     nums1[i] = out[i]
	//   } else{
	//     nums1 = append(nums1, out[i])
	//   }
	// }
}
func trial(a []int) {
	for i, v := range a {
		fmt.Println(v, i)
		i--
	}
}
func searchRange(a []int, t int) []int {
	out := []int{-1, -1}
	mid := len(a) / 2
	start, end := 0, len(a)-1
	for start <= end {
		mid = (start + end) / 2
		if a[mid] == t && (mid == 0 || a[mid] > a[mid-1]) {
			out[0] = mid
		}
		if a[mid] > t {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	mid = len(a) / 2
	start, end = 0, len(a)-1
	for start <= end {
		mid = (start + end) / 2
		if a[mid] == t && (mid == len(a)-1 || a[mid] < a[mid+1]) {
			out[1] = mid
		}
		if a[mid] > t {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return out
}

func binarySearch(a []int, t int) int {
	mid := len(a) / 2
	start, end := 0, len(a)-1
	for start <= end {
		mid = (start + end) / 2
		if a[mid] == t {
			return mid
		}
		if a[mid] > t {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

func painter(a []int, x int, time int) int {
	sum := 0
	count := 1
	for _, v := range a {
		sum += v
		if sum > time {
			count++
			sum = v
		}
	}
	return count
}

func subUnsort(A []int) []int {
	n := len(A)

	// Initialize the left and right indices to -1, indicating no subarray found.
	left, right := -1, -1

	// Traverse the array from left to right to find the first element that is out of place.
	for i := 1; i < n; i++ {
		if A[i] < A[i-1] {
			left = i - 1
			break
		}
	}

	// If no left index is found, the array is already sorted.
	if left == -1 {
		return []int{-1}
	}

	// Traverse the array from right to left to find the first element that is out of place.
	for j := n - 2; j >= 0; j-- {
		if A[j] > A[j+1] {
			right = j + 1
			break
		}
	}

	// Find the minimum and maximum values within the subarray.
	minVal, maxVal := A[left], A[left]
	for i := left; i <= right; i++ {
		if A[i] < minVal {
			minVal = A[i]
		}
		if A[i] > maxVal {
			maxVal = A[i]
		}
	}

	// Extend the left subarray to the left as long as the elements are greater than the minimum value.
	for left > 0 && A[left-1] > minVal {
		left--
	}

	// Extend the right subarray to the right as long as the elements are less than the maximum value.
	for right < n-1 && A[right+1] < maxVal {
		right++
	}

	// Return the left and right indices of the subarray.
	return []int{left, right}
}

func nearestPoint(points [][]int, k int) [][]int {
	comparer := func(p1, p2 []int) bool {
		distanceP1 := p1[0]*p1[0] + p1[1]*p1[1]
		distanceP2 := p2[0]*p2[0] + p2[1]*p2[1]
		return distanceP1 < distanceP2
	}

	sort.SliceStable(points, func(i, j int) bool {
		return comparer(points[i], points[j])
	})
	return points[:k]

}

func countSort(a []int) []int {
	q := maxOfSlice(a)
	n := len(a)
	b := []int{}
	c := make([]int, q+1)
	for i := 0; i < n; i++ {
		c[a[i]]++
	}
	for i := 0; i < len(c); i++ {
		for c[i] > 0 {
			b = append(b, i)
			c[i]--
		}
	}
	return b
}

func maxOfSlice(a []int) int {
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}
