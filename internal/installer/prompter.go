package installer

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Prompter abstracts user interaction for testability.
type Prompter interface {
	// Confirm asks the user a yes/no question and returns true for yes.
	Confirm(message string) bool
}

// StdinPrompter is the production implementation that reads from stdin.
type StdinPrompter struct {
	// Input is the reader to use. Defaults to os.Stdin if nil.
	Input io.Reader
}

// Confirm implements Prompter by reading from Input (or os.Stdin).
func (p *StdinPrompter) Confirm(message string) bool {
	fmt.Printf("%s [y/n]: ", message)
	input := p.Input
	if input == nil {
		input = os.Stdin
	}
	reader := bufio.NewReader(input)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(strings.ToLower(response))
	return response == "y" || response == "yes"
}

// MockPrompter is a test implementation with configurable responses.
type MockPrompter struct {
	// Response is returned by all Confirm calls.
	Response bool
	// Calls records all messages passed to Confirm.
	Calls []string
}

// Confirm implements Prompter for testing.
func (m *MockPrompter) Confirm(message string) bool {
	m.Calls = append(m.Calls, message)
	return m.Response
}
