package main

import "fmt"

// PathStrategy is the strategy interface
type PathStrategy interface {
	CalculatePath(start, destination string) string
}

// WalkStrategy is a concrete strategy for walking
type WalkStrategy struct{}

func (w *WalkStrategy) CalculatePath(start, destination string) string {
	return fmt.Sprintf("Calculating walking path from %s to %s", start, destination)
}

// BikeStrategy is a concrete strategy for biking
type BikeStrategy struct{}

func (b *BikeStrategy) CalculatePath(start, destination string) string {
	return fmt.Sprintf("Calculating biking path from %s to %s", start, destination)
}

// CarStrategy is a concrete strategy for driving
type CarStrategy struct{}

func (c *CarStrategy) CalculatePath(start, destination string) string {
	return fmt.Sprintf("Calculating driving path from %s to %s", start, destination)
}

// PathStrategyFactory is a factory for creating path strategies
type PathStrategyFactory struct{}

func (f *PathStrategyFactory) GetStrategy(mode string) (PathStrategy, error) {
	switch mode {
	case "walk":
		return &WalkStrategy{}, nil
	case "bike":
		return &BikeStrategy{}, nil
	case "car":
		return &CarStrategy{}, nil
	default:
		return nil, fmt.Errorf("invalid mode of transportation: %s", mode)
	}
}

// Navigator is the context
type Navigator struct {
	factory *PathStrategyFactory
}

// NewNavigator creates a new Navigator
func NewNavigator(factory *PathStrategyFactory) *Navigator {
	return &Navigator{factory: factory}
}

// FindPath gets the strategy based on the mode and calculates the path
func (n *Navigator) FindPath(from, to, mode string) (string, error) {
	strategy, err := n.factory.GetStrategy(mode)
	if err != nil {
		return "", err
	}
	return strategy.CalculatePath(from, to), nil
}

func mainFactory() {
	// Create the factory
	factory := &PathStrategyFactory{}
	// Create the navigator
	navigator := NewNavigator(factory)

	// Calculate path using walk mode
	path, err := navigator.FindPath("Home", "Park", "walk")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(path)
	}

	// Calculate path using bike mode
	path, err = navigator.FindPath("Home", "Park", "bike")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(path)
	}

	// Calculate path using car mode
	path, err = navigator.FindPath("Home", "Park", "car")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(path)
	}

	// Handle invalid mode
	path, err = navigator.FindPath("Home", "Park", "fly")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(path)
	}
}

// Behavioural Design Patterns
// The Strategy Pattern allows you to define a family of algorithms,
// encapsulate each one, and make them interchangeable. This pattern
//  lets the algorithm vary independently from the clients that use it.

// •	PathStrategy Interface: Defines the method that all concrete strategies must implement.
// •	Concrete Strategies: WalkStrategy, BikeStrategy, and CarStrategy each implement the PathStrategy interface in their own way.
// •	PathStrategyFactory: Creates the appropriate strategy based on the mode.
// •	Navigator (Context): Uses the factory to get the appropriate strategy based on the mode and calculates the path using the selected strategy.
// •	Usage: In the main function, different path strategies are obtained from the factory based on the mode, and the path is calculated accordingly.
