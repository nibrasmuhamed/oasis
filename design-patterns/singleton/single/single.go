package single

import (
	"sync"
)

// Define the struct
type dbc struct {
	// Add your fields here
}

var instance *dbc
var once sync.Once

// GetInstance returns the singleton instance of dbc
func GetInstance() *dbc {
	once.Do(func() {
		instance = &dbc{}
	})
	return instance
}
