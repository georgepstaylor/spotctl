package pager

import (
	"bytes"
	"os"
	"testing"
)

func TestPager_Write_ShortContent(t *testing.T) {
	// Test that short content doesn't use pager
	p := NewPager()
	p.MinLines = 10 // Set minimum to 10 lines

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	shortContent := "Line 1\nLine 2\nLine 3\n"
	err := p.Write(shortContent)

	w.Close()
	os.Stdout = oldStdout

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Read the captured output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if output != shortContent {
		t.Errorf("Expected output %q, got %q", shortContent, output)
	}
}

func TestPager_Write_Disabled(t *testing.T) {
	// Test that disabled pager always writes directly
	p := NewPager()
	p.Disable = true

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	content := "Line 1\nLine 2\nLine 3\nLine 4\nLine 5\nLine 6\nLine 7\nLine 8\nLine 9\nLine 10\nLine 11\nLine 12\n"
	err := p.Write(content)

	w.Close()
	os.Stdout = oldStdout

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Read the captured output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if output != content {
		t.Errorf("Expected output %q, got %q", content, output)
	}
}

func TestGetPagerCommand(t *testing.T) {
	// Save original environment
	originalPager := os.Getenv("PAGER")
	defer os.Setenv("PAGER", originalPager)

	// Test with PAGER environment variable set to an available command
	os.Setenv("PAGER", "cat") // cat should be available on most systems
	pager, available := getPagerCommand()
	if !available {
		t.Error("Expected pager to be available when PAGER is set to 'cat'")
	}
	if pager != "cat" {
		t.Errorf("Expected 'cat', got %q", pager)
	}

	// Test with PAGER set to non-existent command
	os.Setenv("PAGER", "non-existent-pager-command")
	pager, available = getPagerCommand()
	if available {
		t.Error("Expected pager to be unavailable when set to non-existent command")
	}
	if pager != "" {
		t.Errorf("Expected empty string when pager is unavailable, got %q", pager)
	}

	// Test with no PAGER environment variable
	os.Unsetenv("PAGER")
	pager, available = getPagerCommand()
	if !available {
		t.Error("Expected available=true when no PAGER is set")
	}
	if pager != "" {
		t.Errorf("Expected empty string when no PAGER is set, got %q", pager)
	}
}

func TestIsTerminal(t *testing.T) {
	// This test just ensures the function doesn't panic
	// The actual return value depends on how tests are run
	result := isTerminal()
	t.Logf("isTerminal() returned: %v", result)
}

func TestGetTerminalHeight(t *testing.T) {
	// This test just ensures the function doesn't panic
	// The actual return value depends on the environment
	height := getTerminalHeight()
	t.Logf("getTerminalHeight() returned: %d", height)

	if height < 0 {
		t.Error("Terminal height should not be negative")
	}
}
