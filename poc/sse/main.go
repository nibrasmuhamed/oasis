package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"unsafe"
)

type Product struct {
	InStock  bool    // 1 byte
	Cost     float64 // 8 bytes
	Quantity int32   // 4 bytes
}

// Optimized struct - minimum padding
type OptimizedProduct struct {
	Cost     float64 // 8 bytes
	Quantity int32   // 4 bytes
	InStock  bool    // 1 byte
}

func main() {
	product := Product{}
	optimizedProduct := OptimizedProduct{}
	//log here
	fmt.Printf("size of product struct: %d bytes\n", unsafe.Sizeof(product))
	// size of product struct: 24 bytes

	fmt.Printf("size of optimized product struct: %d bytes\n", unsafe.Sizeof(optimizedProduct))
	// size of optimized product struct: 16 bytes
}
func mainSignals() {
	// Create a channel to receive OS signals
	pid := os.Getpid()
	fmt.Printf("PID: %d\n", pid)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// Create a channel to wait for program termination
	done := make(chan bool, 1)

	// Create a goroutine to handle signals
	go func() {
		// Wait for a signal
		sig := <-sigCh
		fmt.Printf("Received signal: %v\n", sig)

		// Handle the signal
		switch sig {
		case syscall.SIGINT:
			fmt.Println("Received SIGINT. Exiting gracefully...")
		case syscall.SIGTERM:
			fmt.Println("Received SIGTERM. Exiting gracefully...")
		case syscall.SIGQUIT:
			fmt.Println("SIGQUIT")
		}

		// Notify main goroutine that it can exit
		done <- true
	}()

	// Your application logic goes here
	// For example, you can run your server or perform other tasks
	// ...

	// Block the main goroutine until a signal is received
	<-done
}

func sseClient() {
	url := "http://localhost:8080/events?watch=true"

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Perform the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Check if the response status code is OK
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected status code:", resp.Status)
		return
	}

	// Create a new scanner to read from the response body
	scanner := bufio.NewScanner(resp.Body)

	// Read lines from the response body
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Received event:", line)
	}

	// Check for errors from scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading response:", err)
	}
}
func sseServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("sent request to /events?watch=true"))
	})

	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming not supported", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		watch := r.URL.Query().Get("watch")
		if watch != "true" {
			http.Error(w, "Watching not requested", http.StatusBadRequest)
			return
		}

		for {
			// Simulate pushing data to the client every 2 seconds
			fmt.Fprintf(w, "data: %s\n\n", time.Now().Format(time.RFC3339))
			flusher.Flush()
			time.Sleep(2 * time.Second)
		}
	})

	serverAddr := "localhost:8080"
	fmt.Printf("Server listening on http://%s\n", serverAddr)
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

type Response struct {
	Type   string `json:"type"`
	Object struct {
		Kind       string `json:"kind"`
		APIVersion string `json:"apiVersion"`
		Metadata   struct {
			Name              string    `json:"name"`
			Namespace         string    `json:"namespace"`
			UID               string    `json:"uid"`
			ResourceVersion   string    `json:"resourceVersion"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Annotations       struct {
				EndpointsKubernetesIoLastChangeTriggerTime time.Time `json:"endpoints.kubernetes.io/last-change-trigger-time"`
			} `json:"annotations"`
		} `json:"metadata"`
		Subsets []struct {
			Addresses []struct {
				IP        string `json:"ip"`
				NodeName  string `json:"nodeName"`
				TargetRef struct {
					Kind      string `json:"kind"`
					Namespace string `json:"namespace"`
					Name      string `json:"name"`
					UID       string `json:"uid"`
				} `json:"targetRef"`
			} `json:"addresses"`
			Ports []struct {
				Port     int    `json:"port"`
				Protocol string `json:"protocol"`
			} `json:"ports"`
		} `json:"subsets"`
	} `json:"object"`
}

func k8sdata() {
	certPath := "./client-cert.pem"
	keyPath := "./client-key.pem"

	url := "https://192.168.72.2:16443/api/v1/watch/namespaces/default/endpoints?fieldSelector=metadata.name=usage-engine-service"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		fmt.Println("Error loading client certificate and key:", err)
		return
	}

	caCert, err := ioutil.ReadFile(certPath)
	if err != nil {
		log.Fatalf("Error opening cert file %s, Error: %s", certPath, err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	t := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{cert},
			RootCAs:            caCertPool,
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Timeout:   0,
		Transport: t,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected status code:", resp.StatusCode)
		return
	}

	decoder := json.NewDecoder(resp.Body)
	for {
		var data Response
		if err := decoder.Decode(&data); err == nil {
			fmt.Printf("Received event: %+v\n", data)
		} else {
			break
		}
	}
}

func getTLSConfig(host, caCertFile string, certOpt tls.ClientAuthType) *tls.Config {
	var caCert []byte
	var err error
	var caCertPool *x509.CertPool
	if certOpt > tls.RequestClientCert {
		caCert, err = ioutil.ReadFile(caCertFile)
		if err != nil {
			log.Fatal("Error opening cert file", caCertFile, ", error ", err)
		}
		caCertPool = x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
	}

	return &tls.Config{
		ServerName: host,
		ClientAuth: certOpt,
		ClientCAs:  caCertPool,
		MinVersion: tls.VersionTLS12, // TLS versions below 1.2 are considered insecure - see https://www.rfc-editor.org/rfc/rfc7525.txt for details
	}
}

type SessionCounter struct {
	mu    sync.Mutex
	count int
}

func (sc *SessionCounter) GetUniqueValue(sessionID string) int {
	sum := 0
	for _, v := range sessionID {
		sum += int(v)
	}
	// Use the current timestamp and a counter to generate a unique value
	uniqueValue := ((int(time.Now().UnixMilli())) - times) + sum

	return uniqueValue
}

// Simple hash function for string
func hashString(s string) int {

	const maxInt = int(^uint(0) >> 1)
	hash := 0
	for _, c := range s {
		hash = (31*hash + int(c)) % maxInt
	}
	return hash
}

var times int

func mains() {
	// k8sdata()
	maxInt32 := int32(2147483647)

	// Incrementing the maximum value
	maxInt32++

	// The result will be the minimum representable value for int32
	fmt.Println(maxInt32) // Output: -2147483648
	counter := &SessionCounter{}
	timeNow := time.Now()
	timeFormatted := timeNow.Format("2006-01-02T15:04:05Z")
	fmt.Println(timeFormatted)
	ss := math.MaxInt32
	times = int(time.Now().UnixMilli())
	fmt.Println("Maximum representable integer value:", ss)
	fmt.Println(int(time.Now().UnixMilli()))
	// Test with multiple goroutines
	// var wg sync.WaitGroup
	// for i := 1; i <= 5; i++ {
	// 	wg.Add(1)
	// 	go func(sessionID string) {
	// 		defer wg.Done()

	// 		// Simulate requests with unique session IDs
	// 		for j := 1; j <= 3; j++ {
	// 			uniqueValue := counter.GetUniqueValue(sessionID)
	// 			fmt.Printf("Session ID: %s, Unique Value: %d\n", sessionID, uniqueValue)
	// 			time.Sleep(time.Millisecond * 100) // Simulate some processing time
	// 		}
	// 	}(fmt.Sprintf("Session%d", i))
	// }
	for j := 1; j <= 100; j++ {
		uniqueValue := counter.GetUniqueValue("s")
		fmt.Printf("Session ID: %d, Unique Value: %d\n", j, uniqueValue)
		time.Sleep(time.Millisecond * 100) // Simulate some processing time
	}
	// wg.Wait()
	//  2,147,483,647
	// 2147483647
	// 1705589934725
}
