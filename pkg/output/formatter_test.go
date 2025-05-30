package output

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

// Test types that mimic the client types
type TestRegion struct {
	APIVersion string       `json:"apiVersion,omitempty"`
	Kind       string       `json:"kind,omitempty"`
	Metadata   TestMetadata `json:"metadata,omitempty"`
	Spec       TestSpec     `json:"spec,omitempty"`
}

type TestMetadata struct {
	Name string `json:"name,omitempty"`
	UID  string `json:"uid,omitempty"`
}

type TestSpec struct {
	Country     string       `json:"country,omitempty"`
	Description string       `json:"description,omitempty"`
	Provider    TestProvider `json:"provider,omitempty"`
}

type TestProvider struct {
	ProviderType       string `json:"providerType,omitempty"`
	ProviderRegionName string `json:"providerRegionName,omitempty"`
}

type TestRegionList struct {
	APIVersion string       `json:"apiVersion,omitempty"`
	Items      []TestRegion `json:"items"`
	Kind       string       `json:"kind,omitempty"`
}

func TestFormatter_OutputJSON(t *testing.T) {
	data := &TestRegionList{
		APIVersion: "v1",
		Kind:       "RegionList",
		Items: []TestRegion{
			{
				APIVersion: "v1",
				Kind:       "Region",
				Metadata: TestMetadata{
					Name: "us-east-1",
				},
				Spec: TestSpec{
					Country: "United States",
				},
			},
		},
	}

	formatter := NewFormatter(OutputOptions{Format: JSONFormat})

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := formatter.Output(data, nil)
	if err != nil {
		t.Fatalf("Output failed: %v", err)
	}

	// Restore stdout and read output
	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Verify JSON output
	if !strings.Contains(output, `"apiVersion": "v1"`) {
		t.Error("JSON output should contain apiVersion")
	}
	if !strings.Contains(output, `"us-east-1"`) {
		t.Error("JSON output should contain region name")
	}
}

func TestFormatter_OutputTable(t *testing.T) {
	data := &TestRegionList{
		Items: []TestRegion{
			{
				Metadata: TestMetadata{
					Name: "us-east-1",
					UID:  "12345",
				},
				Spec: TestSpec{
					Country:     "United States",
					Description: "US East Region",
					Provider: TestProvider{
						ProviderType:       "aws",
						ProviderRegionName: "us-east-1",
					},
				},
			},
			{
				Metadata: TestMetadata{
					Name: "empty-region",
				},
				Spec: TestSpec{
					Country: "", // Test empty values
				},
			},
		},
	}

	config := &TableConfig{
		Columns: []TableColumn{
			{Header: "NAME", Field: "metadata.name", Default: "N/A"},
			{Header: "COUNTRY", Field: "spec.country", Default: "N/A"},
			{Header: "PROVIDER", Field: "spec.provider.providerType", Default: "N/A"},
		},
		DetailCols: []TableColumn{
			{Header: "DESCRIPTION", Field: "spec.description", Default: "N/A"},
		},
		WideCols: []TableColumn{
			{Header: "UID", Field: "metadata.uid", Default: "N/A"},
		},
	}

	// Test basic table
	formatter := NewFormatter(OutputOptions{Format: TableFormat})

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := formatter.Output(data, config)
	if err != nil {
		t.Fatalf("Output failed: %v", err)
	}

	// Restore stdout and read output
	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Verify table output
	if !strings.Contains(output, "us-east-1") {
		t.Error("Table should contain region name")
	}
	if !strings.Contains(output, "United States") {
		t.Error("Table should contain country")
	}
	if !strings.Contains(output, "N/A") {
		t.Error("Table should contain 'N/A' for empty fields")
	}
	if !strings.Contains(output, "NAME") {
		t.Error("Table should contain headers")
	}
	if strings.Contains(output, "DESCRIPTION") {
		t.Error("Basic table should not contain detail columns")
	}
}

func TestFormatter_OutputTableDetails(t *testing.T) {
	data := &TestRegionList{
		Items: []TestRegion{
			{
				Metadata: TestMetadata{Name: "test-region"},
				Spec: TestSpec{
					Country:     "Test Country",
					Description: "Test Description",
				},
			},
		},
	}

	config := &TableConfig{
		Columns: []TableColumn{
			{Header: "NAME", Field: "metadata.name"},
		},
		DetailCols: []TableColumn{
			{Header: "DESCRIPTION", Field: "spec.description"},
		},
	}

	formatter := NewFormatter(OutputOptions{
		Format:      TableFormat,
		ShowDetails: true,
	})

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := formatter.Output(data, config)
	if err != nil {
		t.Fatalf("Output failed: %v", err)
	}

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "DESCRIPTION") {
		t.Error("Detailed table should contain detail columns")
	}
	if !strings.Contains(output, "Test Description") {
		t.Error("Detailed table should contain description value")
	}
}

func TestFormatter_GetFieldValue(t *testing.T) {
	formatter := NewFormatter(OutputOptions{})

	item := TestRegion{
		Metadata: TestMetadata{
			Name: "test-region",
		},
		Spec: TestSpec{
			Provider: TestProvider{
				ProviderType: "aws",
			},
		},
	}

	tests := []struct {
		field    string
		expected string
	}{
		{"metadata.name", "test-region"},
		{"spec.provider.providerType", "aws"},
		{"nonexistent.field", ""},
		{"", ""},
	}

	for _, test := range tests {
		result := formatter.getFieldValue(item, test.field)
		if result != test.expected {
			t.Errorf("getFieldValue(%q) = %q, expected %q", test.field, result, test.expected)
		}
	}
}
