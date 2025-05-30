package pager

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"golang.org/x/term"
)

// max returns the larger of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Pager handles automatic paging of long output
type Pager struct {
	// ForceEnable forces pager even when output is not a terminal
	ForceEnable bool
	// Disable completely disables the pager
	Disable bool
	// MinLines is the minimum number of lines before pager is used
	MinLines int
}

// NewPager creates a new pager with default settings
func NewPager() *Pager {
	return &Pager{
		MinLines: 10, // Default to paging if output is more than 10 lines
	}
}

// Write outputs the content, using a pager if appropriate
func (p *Pager) Write(content string) error {
	// If pager is disabled, just write to stdout
	if p.Disable {
		fmt.Print(content)
		return nil
	}

	// If output is not a terminal and not forced, write directly
	if !p.ForceEnable && !isTerminal() {
		fmt.Print(content)
		return nil
	}

	// Count lines in content
	lines := strings.Count(content, "\n")
	if content != "" && !strings.HasSuffix(content, "\n") {
		lines++ // Count last line if it doesn't end with newline
	}

	// If content is short enough, write directly
	termHeight := getTerminalHeight()

	// Use terminal height if available, otherwise fall back to MinLines
	threshold := p.MinLines
	if termHeight > 0 {
		// Use 80% of terminal height as threshold, but at least MinLines
		threshold = max(termHeight*4/5, p.MinLines)
	}

	if lines <= threshold {
		fmt.Print(content)
		return nil
	}

	// Use pager for long content
	return p.usePager(content)
}

// usePager pipes content through a pager program
func (p *Pager) usePager(content string) error {
	pagerCmd, available := getPagerCommand()
	if pagerCmd == "" {
		if !available {
			// User's configured pager is not available, show warning
			fmt.Fprintf(os.Stderr, "Warning: Configured pager is not available. Output will be displayed directly.\n")
		}
		// No pager available, write directly
		fmt.Print(content)
		return nil
	}

	// Parse pager command (might have flags)
	parts := strings.Fields(pagerCmd)
	if len(parts) == 0 {
		fmt.Print(content)
		return nil
	}

	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Create pipe to send content to pager
	stdin, err := cmd.StdinPipe()
	if err != nil {
		// Fallback to direct output
		fmt.Print(content)
		return nil
	}

	// Start the pager
	if err := cmd.Start(); err != nil {
		stdin.Close()
		fmt.Print(content)
		return nil
	}

	// Write content to pager
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, content)
	}()

	// Wait for pager to finish
	return cmd.Wait()
}

// getPagerCommand returns the pager command to use, respecting user configuration
// Returns (command, available) where:
// - command: the pager command to use (empty if none)
// - available: true if no error occurred, false if user's pager is configured but not available
func getPagerCommand() (string, bool) {
	// Check environment variable first
	if pager := os.Getenv("PAGER"); pager != "" {
		// User has explicitly set a pager - validate it exists
		parts := strings.Fields(pager)
		if len(parts) > 0 {
			if _, err := exec.LookPath(parts[0]); err == nil {
				return pager, true
			}
			// User's pager is configured but not available
			return "", false
		}
	}

	// No pager configured by user - this is fine, output directly
	return "", true
}

// isTerminal checks if stdout is a terminal
func isTerminal() bool {
	return term.IsTerminal(int(os.Stdout.Fd()))
}

// getTerminalHeight returns the terminal height in lines, or 0 if unknown
func getTerminalHeight() int {
	// Try to get terminal size
	if width, height, err := term.GetSize(int(os.Stdout.Fd())); err == nil {
		_ = width // width is available but we only need height
		return height
	}

	// Fallback: try reading from environment
	if lines := os.Getenv("LINES"); lines != "" {
		if h, err := strconv.Atoi(lines); err == nil {
			return h
		}
	}

	// Fallback: try using stty
	if height := getTerminalHeightStty(); height > 0 {
		return height
	}

	return 0
}

// getTerminalHeightStty gets terminal height using stty command
func getTerminalHeightStty() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, err := cmd.Output()
	if err != nil {
		return 0
	}

	parts := strings.Fields(string(bytes.TrimSpace(output)))
	if len(parts) >= 2 {
		if height, err := strconv.Atoi(parts[0]); err == nil {
			return height
		}
	}

	return 0
}

// WriteToWriter is a helper that captures output from a writer function and pages it
func (p *Pager) WriteToWriter(writerFunc func(io.Writer) error) error {
	// If pager is disabled, write directly to stdout
	if p.Disable {
		return writerFunc(os.Stdout)
	}

	// If output is not a terminal and not forced, write directly
	if !p.ForceEnable && !isTerminal() {
		return writerFunc(os.Stdout)
	}

	// Capture output in a buffer
	var buf bytes.Buffer
	if err := writerFunc(&buf); err != nil {
		return err
	}

	// Page the captured content
	return p.Write(buf.String())
}
