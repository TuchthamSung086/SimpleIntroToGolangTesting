package printer_test

import (
	"fmt"
	"testing"
	printer "testing-demo/3_mock"

	"github.com/stretchr/testify/mock"
)

// PrinterMock is a mock object. This object mocks the Printer interface.
type PrinterMock struct {
	mock.Mock
}

func (d *PrinterMock) Print(paperSize string, content string) error {
	// indicates that the function is called
	args := d.Called(paperSize, content)

	fmt.Println("Printing from mock object...")
	return args.Error(0)
}

func TestPrintA4(t *testing.T) {
	// create a printer mock
	pm := &PrinterMock{}

	// set expectation that the mocked function will receive
	// "A4" and "Hello world" as its input and got a return value
	// as nil as its output
	pm.On("Print", "A4", "hello world").Return(nil)

	// create a computer struct
	computer := printer.Computer{Pr: pm}

	// call PrintA4
	// The PrintA4 will call printer.Print
	computer.PrintA4("hello world")

	// assert that our expectations are met
	pm.AssertExpectations(t)
}
