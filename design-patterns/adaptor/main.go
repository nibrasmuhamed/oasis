package main

import "fmt"

// ModernPrinter is the interface expected by the new system
type ModernPrinter interface {
	Print(msg string) string
}

// LegacyPrinter is the existing printer used by the legacy system
type LegacyPrinter struct{}

func (lp *LegacyPrinter) PrintMessage(msg string) string {
	return fmt.Sprintf("Legacy Printer: %s", msg)
}

// PrinterAdapter adapts the LegacyPrinter to the ModernPrinter interface
type PrinterAdapter struct {
	legacyPrinter *LegacyPrinter
}

func (pa *PrinterAdapter) Print(msg string) string {
	return pa.legacyPrinter.PrintMessage(msg)
}

func main() {
	legacyPrinter := &LegacyPrinter{}
	adapter := &PrinterAdapter{legacyPrinter}

	message := adapter.Print("Hello, World!")
	fmt.Println(message)
}

// The Adapter Pattern is a structural design pattern that
// allows objects with incompatible interfaces to work together.
// It acts as a bridge between two incompatible interfaces by
// wrapping an existing class with a new interface.
// adapter is useful when integrating with 3rd party system.
