package printer

import "fmt"

// Printer is an interface which implements the Print function.
type Printer interface {
	Print(paperSize string, content string) error
}

// MyPrinter is a type of printer that could do printing jobs
type MyPrinter struct{}

// Print is a method of type *MyPrinter
// You may give the paperSize and content to it and it will print for you.
func (e *MyPrinter) Print(paperSize string, content string) error {
	// Hard, confusing, database and external parties calling stuff here
	fmt.Println("Printing:", content, ", with paper size:", paperSize)
	return nil
}
