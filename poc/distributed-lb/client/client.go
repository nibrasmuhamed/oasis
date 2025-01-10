package client

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"distributed-lb/hash"
	"distributed-lb/message"

	"github.com/cenkalti/backoff/v4"
)

type Client struct {
	consistent         *hash.Consistent
	httpClient         *http.Client
	backOff            *backoff.ExponentialBackOff
	isConnectionActive bool
	connectionTimeout  time.Duration
	url                string
}

func New(url string) *Client {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.DisableKeepAlives = false
	transport.MaxIdleConns = 10
	transport.MaxConnsPerHost = 0
	transport.MaxIdleConnsPerHost = 0
	transport.IdleConnTimeout = 100 * time.Second
	//fmt.Printf("%+v\n", transport)

	bo := backoff.NewExponentialBackOff()
	bo.InitialInterval = 500 * time.Millisecond
	bo.RandomizationFactor = 0.5
	bo.Multiplier = 1.5
	bo.MaxInterval = 60 * time.Second
	bo.MaxElapsedTime = 1 * time.Hour
	bo.Reset()

	return &Client{
		httpClient: &http.Client{
			Transport: transport,
		},
		backOff:    bo,
		url:        url,
		connectionTimeout: time.Minute,
	}
}

func (client *Client) Run(cancel context.CancelCauseFunc) {
	for {
		err := client.listen()
		if err != nil {
			client.isConnectionActive = false
			log.Println("Error: ", err)
			log.Println("Retrying....")
			timeout := client.backOff.NextBackOff()
			if timeout == backoff.Stop {
				cancel(errors.New("Max Elapsed Time reached to connect to coordinator: " + client.url))
				return
			}
			time.Sleep(timeout)
		}
	}
}

func (client *Client) listen() error {
	// Create a new HTTP request to the SSE server
	request, err := http.NewRequest("GET", client.url, nil)
	if err != nil {
		return err
	}
	// Set the "Accept" header to "text/event-stream" to indicate support for Server-Sent Events
	request.Header.Set("Accept", "text/event-stream")

	// Make the request and check for errors
	response, err := client.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	client.isConnectionActive = true
	// Check if the server supports Server-Sent Events
	if response.Header.Get("Content-Type") != "text/event-stream" {
		log.Fatal("Server does not support Server-Sent Events")
	}
	reader := bufio.NewReader(response.Body)
	for {
		client.backOff.Reset()
		var msg message.Message
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		data := strings.TrimSuffix(line, "\n")
		err = json.Unmarshal([]byte(data), &msg)

		if err != nil {
			return err
		}

		client.consistent = msg.Update(client.consistent)
	}
}

func (client *Client) LocateKey(key []byte) (hash.Member, error) {
	if client.consistent == nil ||
		(!client.isConnectionActive && client.backOff.GetElapsedTime() > client.connectionTimeout) {
		return hash.Member{}, errors.New("Cluster error: Unable to fetch cluster information")
	}
	m := client.consistent.LocateKey(key)
	if m.Name == "" {
		return m, errors.New("Hash Error: Node not found for the key")
	}
	return m, nil
}
