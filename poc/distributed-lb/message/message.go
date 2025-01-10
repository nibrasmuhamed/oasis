package message

import (
	"distributed-lb/hash"
	"log"
)

const (
	ERROR       = -1
	HEALTHCHECK = 0
	INIT        = 1
	ADD         = 2
	REMOVE      = 3
)

type Message struct {
	Command           int
	Error             string
	Members           hash.MemberList
	PartitionCount    int
	ReplicationFactor int
	Load              float64
	Time              string
}

func (msg Message) Update(c *hash.Consistent) *hash.Consistent {
	switch msg.Command {
	case INIT:
		cfg := hash.Config{
			PartitionCount:    msg.PartitionCount,
			ReplicationFactor: msg.ReplicationFactor,
			Load:              msg.Load,
		}
		c = hash.New(msg.Members, cfg)
		log.Printf("Initializing node:  %+v\n", msg)
	case ADD:
		for _, m := range msg.Members {
			//fmt.Println("Adding new member:", m.String())
			c.Add(m)
		}
		log.Printf("Adding node: %+v\n", msg.Members)
	case REMOVE:
		// oldMembers := c.GetMembers()
		// fmt.Println(oldMembers)
		for _, m := range msg.Members {
			//fmt.Println("Removing member:", m.String())
			c.Remove(m.String())
		}
		log.Printf("Deleting node: %+v\n", msg.Members)
	case ERROR:
		log.Println("Error: ", msg.Error)
		return nil
	case HEALTHCHECK:
		log.Println("Received HealthCheck command")
	}
	return c
}
