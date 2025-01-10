package main

import "fmt"

// Prototype is the interface that all prototypes must implement
type Prototype[T any] interface {
	Clone() T
}

// The Prototype Pattern is a creational design pattern that involves
// creating new objects by copying a prototype instance. This pattern
// is particularly useful when the cost of creating a new instance is
// more expensive than copying an existing instance.

// Document represents a concrete prototype
type Document struct {
	Title   string
	Content string
}

// Clone creates a copy of the Document
func (d Document) Clone() Document {
	return Document{
		Title:   d.Title,
		Content: d.Content,
	}
}

func main() {
	// Create an original document
	original := Document{
		Title:   "Original Title",
		Content: "Original Content",
	}

	// Clone the original document
	var prototype Prototype[Document] = original
	clone := prototype.Clone()

	// Print both documents
	fmt.Printf("Original: %+v\n", original)
	fmt.Printf("Clone: %+v\n", clone)

	// Modify the clone
	clone.Title = "Modified Title"
	clone.Content = "Modified Content"

	// Print both documents again to show they are different
	fmt.Printf("After modification\n")
	fmt.Printf("Original: %+v\n", original)
	fmt.Printf("Clone: %+v\n", clone)
}
