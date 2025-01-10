package main

import (
	"context"
	"distributed-lb/client"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var c *client.Client

func main() {
	port := flag.String("port", "1", "port number")
	flag.Parse()
	ctx, cancel := context.WithCancelCause(context.Background())	

	c = client.New("http://127.0.0.1:8081")
	go c.Run(cancel)
	
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", getCustomer)
		err := http.ListenAndServe("127.0.0.1:900"+*port, mux)
		if err != nil {
			fmt.Println("Error starting the http server:" + err.Error())
			cancel(err)
		}
	}()
	<-ctx.Done()
	fmt.Println(context.Cause(ctx))
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/customer/")
	m, err := c.LocateKey([]byte(id))
	if (err != nil ) {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Error fetching the key: "+err.Error())
		return
	}
	fmt.Printf("Id: %s, Member: %s\n", id, m)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Key is in node: "+m.String()+"\n")
}