package coordinator

import (
	"distributed-lb/hash"
	"distributed-lb/message"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	PartitionCount    = 16000
	ReplicationFactor = 1000
	Load              = 1.2
)

type Listeners struct {
	Id      int64
	Message chan message.Message
}

type Coordinator struct {
	listeners          []*Listeners
	mu                 sync.RWMutex
	sequence           int64
	consistent         *hash.Consistent
	config             hash.Config
	healthCheckTimeout time.Duration
	StateFile          string
}

func New(members []hash.Member) *Coordinator {
	coord := Coordinator{
		config: hash.Config{
			PartitionCount:    PartitionCount,
			ReplicationFactor: ReplicationFactor,
			Load:              Load,
		},
		healthCheckTimeout: time.Minute,
		StateFile:          "members.json",
	}
	oldMembers := coord.readPreviousState()
	//fmt.Println("Old Members: ", oldMembers)
	c := hash.New(oldMembers, coord.config)
	oldP := c.GetPartitionList()
	coord.consistent = hash.New(members, coord.config)
	fmt.Println("Delete Partition: ", coord.rePartition(oldP))
	coord.saveState()

	coord.addHttpHandler()
	go coord.healthCheck()

	port := "8081" // TODO: Read it from env
	server := &http.Server{
		Addr: ":" + port,
		//WriteTimeout: time.Second * 60,
	}
	go func() {
		// Start the server on port 8080
		fmt.Println("Server is running on http://localhost:" + port)
		if err := server.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	return &coord
}

func (coord *Coordinator) healthCheck() {
	for {
		m := message.Message{
			Command: message.HEALTHCHECK,
		}
		coord.mu.Lock()
		coord.broadCast(m)
		coord.mu.Unlock()
		time.Sleep(coord.healthCheckTimeout)
	}
}

func (coord *Coordinator) addHttpHandler() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set the Content-Type header to text/event-stream
		w.Header().Set("Content-Type", "text/event-stream")
		// Set the Cache-Control header to prevent caching
		w.Header().Set("Cache-Control", "no-cache")
		// Enable CORS (Cross-Origin Resource Sharing)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Connection", "keep-alive")
		listener := Listeners{
			Message: make(chan message.Message),
		}
		go coord.AddListener(&listener)

		for {
			m := <-listener.Message
			j, _ := json.Marshal(&m)
			_, err := fmt.Fprintf(w, "%s\n", j)
			if err != nil {
				fmt.Println("Client disconnected : " + err.Error())
				break
			}
			w.(http.Flusher).Flush()
		}
		coord.RemoveListener(&listener)
	})
}

func (coord *Coordinator) AddListener(listener *Listeners) {
	coord.mu.Lock()
	defer coord.mu.Unlock()
	coord.sequence++
	listener.Id = coord.sequence
	coord.listeners = append(coord.listeners, listener)
	m := message.Message{
		Command:           message.INIT,
		Time:              time.Now().Format(time.RFC3339),
		Members:           coord.consistent.GetMembers(),
		PartitionCount:    PartitionCount,
		ReplicationFactor: ReplicationFactor,
		Load:              Load,
	}
	listener.Message <- m
	fmt.Printf("Added Listener %d, Total count: %d\n", listener.Id, len(coord.listeners))
}

func (coord *Coordinator) RemoveListener(listener *Listeners) {
	coord.mu.Lock()
	defer coord.mu.Unlock()

	for i, c := range coord.listeners {
		if c.Id == listener.Id {
			close(coord.listeners[i].Message)
			coord.listeners = append(coord.listeners[:i], coord.listeners[i+1:]...)
			fmt.Printf("Removed Listener %d, Total count: %d\n", listener.Id, len(coord.listeners))
			break
		}
	}
}

func (coord *Coordinator) broadCast(message message.Message) {
	for _, c := range coord.listeners {
		c.Message <- message
	}
}

func (coord *Coordinator) rePartition(old map[int]*hash.Member) map[string][]int {
	new := coord.consistent.GetPartitionList()
	// for h := range old {
	// 	fmt.Println(h, ": ", old[h].Name, ",", new[h].Name)
	// }

	delete := map[string][]int{}
	for h := range old {
		n := old[h].Name
		if m, exists := new[h]; exists &&
			m.Name != n && coord.consistent.MemberExists(n) {
			delete[n] = append(delete[n], h)
		}
	}
	return delete
}

// func (coord *Coordinator) updateMembers(delete map[string][]int) {
// 	for hashes, name := range delete {
// 		// send data to nodes
// 		// if node is not reachable, then we need to add it
// 		// potential dead list, if we don't see the node from
//      // k8s service then we delete it until then we keep trying
//      // the node in exponential
// 	}

// }

func (coord *Coordinator) RemoveMember(m hash.Member) {
	coord.mu.Lock()
	defer coord.mu.Unlock()
	coord.consistent.Remove(m.Name)
	fmt.Println("Removing Node: ", m.Name)
	msg := message.Message{
		Command: message.REMOVE,
		Members: []hash.Member{m},
	}
	coord.broadCast(msg)
	coord.saveState()
}

func (coord *Coordinator) AddMember(members []hash.Member) {
	coord.mu.Lock()
	defer coord.mu.Unlock()
	oldPartition := coord.consistent.GetPartitionList()
	for _, m := range members {
		coord.consistent.Add(m)
		fmt.Println("Adding Node: ", m.Name)
	}
	fmt.Println("Remove Partitions", coord.rePartition(oldPartition))
	m := message.Message{
		Command: message.ADD,
		Members: members}
	coord.broadCast(m)
	coord.saveState()
}

func (coord *Coordinator) saveState() {
	file, _ := json.Marshal(coord.consistent.GetMembers())
	err := os.WriteFile(coord.StateFile, file, 0644)
	if err != nil {
		panic(err)
	}
}

func (coord *Coordinator) readPreviousState() []hash.Member {
	var members []hash.Member
	fi, err := os.Stat(coord.StateFile)
	if err != nil {
		panic(err)
	}

	size := fi.Size()
	if size == 0 {
		return members
	}
	data, err := os.ReadFile(coord.StateFile)
	if err != nil {
		panic(err)
	}

	if data != nil {
		err = json.Unmarshal(data, &members)
	}
	if err != nil {
		fmt.Println(err)
	}
	return members
}

func compareLists(oldList, newList []hash.Member) ([]hash.Member, []hash.Member) {
	deleted := make([]hash.Member, 0)
	added := make([]hash.Member, 0)

	// Create a map to store the presence of each item in the old list
	oldMap := make(map[string]bool)
	for _, item := range oldList {
		oldMap[item.Name] = true
	}

	// Check each item in the new list
	for _, item := range newList {
		// If the item is not present in the old list, it's added
		if _, exists := oldMap[item.Name]; !exists {
			added = append(added, item)
		} else {
			// If the item is present in the old list, remove it from the map
			delete(oldMap, item.Name)
		}
	}

	// Remaining items in the oldMap are deleted
	for _, item := range oldList {
		if _, exists := oldMap[item.Name]; exists {
			deleted = append(deleted, item)
		}
	}

	return deleted, added
}

func (coord *Coordinator) GetMembers() []hash.Member {
	return coord.consistent.GetMembers()
}
