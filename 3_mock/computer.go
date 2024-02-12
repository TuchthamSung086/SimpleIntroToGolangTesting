package printer

type Computer struct {
	// Pr is a field of type Printer interface
	Pr Printer
}

// PrintA4 prints content in an A4 paper format
func (c *Computer) PrintA4(content string) error {
	c.Pr.Print("A4", content)
	return nil
}
