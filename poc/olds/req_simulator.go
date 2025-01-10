package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/goccy/go-json"
)

var body ChargingDataRequest
var counter uint64
var client = &http.Client{
	Timeout: 10 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
	},
}

const (
	// Total number of requests to simulate
	concurrency      = 8
	requestPerThread = 500

	numRequests = concurrency * requestPerThread * 3 // Number of concurrent requests
	baseUrl     = "http://localhost:8081"            // Replace with your API endpoint
	initiate    = "/chargingData"
	update      = "/Update"
	terminate   = "/Release"
)

func requestThread(i, j int, q chan int, body ChargingDataRequest) {
	// defer wg.Done()
	for k := i; k < j; k++ {
		body.SubscriberIdentifier = "sub" + fmt.Sprintf("%d", k+1)
		body.MnSConsumerIdentifier = "cust" + fmt.Sprintf("%d", k+1)
		body.InvocationSequenceNumber = 0
		body.MultipleUnitUsage[0].UsedUnitContainer[0].ServiceSpecificUnits = 0
		f, s, err := request(body, baseUrl+initiate)
		atomic.AddUint64(&counter, 1)
		if err != nil || f != 200 {
			fmt.Println("Error making request:", err)
			q <- 0
			continue
		}
		body.MultipleUnitUsage[0].UsedUnitContainer[0].ServiceSpecificUnits = 1
		body.InvocationSequenceNumber = 1
		body.InvocationTimeStamp = body.InvocationTimeStamp.Add(time.Second * 1)
		_, _, err = request(body, baseUrl+"/"+s+update)
		atomic.AddUint64(&counter, 1)

		if err != nil {
			fmt.Println("Error making request:", err)
			q <- 0
			continue
		}
		body.InvocationSequenceNumber = 2
		body.InvocationTimeStamp = body.InvocationTimeStamp.Add(time.Second * 1)
		body.MultipleUnitUsage[0].UsedUnitContainer[0].ServiceSpecificUnits = 1
		_, _, err = request(body, baseUrl+"/"+s+terminate)
		atomic.AddUint64(&counter, 1)
		body.MultipleUnitUsage[0].UsedUnitContainer[0].ServiceSpecificUnits = 0

		if err != nil {
			fmt.Println("Error making request:", err)
			q <- 0
			continue
		}
		q <- 1
	}
}

func sim() {
	success := make(chan int, numRequests)
	fmt.Printf("Simulating %d requests with %d concurrency...\n", numRequests, concurrency)
	startTime := time.Now()
	for i := 0; i < concurrency; i++ {
		go requestThread(i*requestPerThread, (i+1)*requestPerThread, success, body)
	}
	successfulRequests := 0
	failureCall := 0
	for result := range success {
		if result == 1 {
			successfulRequests++
		} else {
			failureCall++
		}
		if failureCall+successfulRequests == concurrency*requestPerThread {
			break
		}
	}
	close(success)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Simulation complete!\n")
	fmt.Printf("Successful Requests: %d\n", counter-uint64(failureCall))
	fmt.Printf("Total Requests: %d\n", counter)
	fmt.Printf("Concurrency: %d\n", concurrency)
	fmt.Printf("Elapsed Time: %v\n", elapsedTime)
	fmt.Printf("TPS: %v\n", counter/uint64(elapsedTime.Seconds()))
}

func init() {
	jsonData, err := ioutil.ReadFile("request.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}
	json.Unmarshal(jsonData, &body)
}

func request(body interface{}, url string) (int, string, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return 0, "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return 0, "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()

	session := resp.Header.Get("X-Session-Id")
	return resp.StatusCode, session, nil
}
