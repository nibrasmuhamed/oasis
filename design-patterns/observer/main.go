// The Observer Pattern is a behavioral design pattern where an object
// (called the subject) maintains a list of its dependents (called observers)
//
//	and notifies them of any state changes, typically by calling one of
//	their methods. This pattern is useful when multiple objects need to be
//	updated whenever the state of one object changes.
package main

import "fmt"

// Observer interface
type Observer interface {
	Update(data string)
}

// Subject interface
type Subject interface {
	Register(observer Observer)
	Deregister(observer Observer)
	NotifyAll()
}

// Concrete Subject
type NewsAgency struct {
	observers []Observer
	news      string
}

// Register adds an observer to the list
func (na *NewsAgency) Register(observer Observer) {
	na.observers = append(na.observers, observer)
}

// Deregister removes an observer from the list
func (na *NewsAgency) Deregister(observer Observer) {
	for i, o := range na.observers {
		if o == observer {
			na.observers = append(na.observers[:i], na.observers[i+1:]...)
			break
		}
	}
}

// NotifyAll sends the latest news to all registered observers
func (na *NewsAgency) NotifyAll() {
	for _, observer := range na.observers {
		observer.Update(na.news)
	}
}

// PublishNews updates the state (news) and notifies all observers
func (na *NewsAgency) PublishNews(news string) {
	na.news = news
	na.NotifyAll()
}

// Concrete Observer
type Reader struct {
	name string
}

// Update receives a notification from the subject
func (r *Reader) Update(news string) {
	fmt.Printf("%s received news: %s\n", r.name, news)
}
func main() {
	// Create a news agency (subject)
	agency := &NewsAgency{}

	// Create some readers (observers)
	reader1 := &Reader{name: "Alice"}
	reader2 := &Reader{name: "Bob"}

	// Register readers
	agency.Register(reader1)
	agency.Register(reader2)

	// Publish news, all readers get notified
	agency.PublishNews("Golang 1.20 Released!")

	// Deregister one reader and publish news again
	agency.Deregister(reader1)
	agency.PublishNews("New Go Modules Tutorial Available!")
}
