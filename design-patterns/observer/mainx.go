package main

import "fmt"

type Observerx interface {
	Update(string)
}

type Producer struct {
	observers []Observer
}

type Consumer struct {
}

func (p *Producer) Attach(o Observer) {
	p.observers = append(p.observers, o)
}
func NewProducer() *Producer {
	return &Producer{}
}
func (p *Producer) Notify(msg string) {
	for _, o := range p.observers {
		o.Update(msg)
	}
}
func (c *Consumer) Update(msg string) {
	fmt.Println(msg)
}

// func main() {
// 	p := NewProducer()
// 	c := &Consumer{}
// 	p.Attach(c)
// 	p.Notify("Hello World")
// }
