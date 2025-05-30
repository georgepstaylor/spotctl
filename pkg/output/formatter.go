package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"

	"gopkg.in/yaml.v3"

	"github.com/georgetaylor/spotctl/pkg/pager"
)

// OutputFormat represents supported output formats
type OutputFormat string

const (
	TableFormat OutputFormat = "table"
	JSONFormat  OutputFormat = "json"
	YAMLFormat  OutputFormat = "yaml"
)

// OutputOptions contains options for formatting output
type OutputOptions struct {
	Format      OutputFormat
	ShowDetails bool
	WideOutput  bool
}

// TableColumn represents a column in table output
type TableColumn struct {
	Header  string
	Field   string // JSON path-like field selector
	Width   int    // Optional fixed width
	Default string // Default value for empty fields
}

// TableConfig defines how to render a resource as a table
type TableConfig struct {
	Columns    []TableColumn
	DetailCols []TableColumn // Additional columns for --details
	WideCols   []TableColumn // Additional columns for --wide
}

// Formatter handles output formatting for CLI commands
type Formatter struct {
	options OutputOptions
	pager   *pager.Pager
}

// NewFormatter creates a new formatter with the given options
func NewFormatter(options OutputOptions) *Formatter {
	return &Formatter{
		options: options,
		pager:   pager.NewPager(),
	}
}

// NewFormatterWithPager creates a new formatter with custom pager settings
func NewFormatterWithPager(options OutputOptions, p *pager.Pager) *Formatter {
	return &Formatter{
		options: options,
		pager:   p,
	}
}

// Output formats and outputs the given data according to the formatter's options
// This method uses the pager for long output
func (f *Formatter) Output(data interface{}, tableConfig *TableConfig) error {
	return f.pager.WriteToWriter(func(w io.Writer) error {
		return f.OutputToWriter(w, data, tableConfig)
	})
}

// OutputToWriter formats and outputs data to the specified writer
func (f *Formatter) OutputToWriter(w io.Writer, data interface{}, tableConfig *TableConfig) error {
	switch f.options.Format {
	case JSONFormat:
		return f.outputJSONToWriter(w, data)
	case YAMLFormat:
		return f.outputYAMLToWriter(w, data)
	case TableFormat:
		if tableConfig == nil {
			return fmt.Errorf("table configuration required for table output")
		}
		return f.outputTableToWriter(w, data, tableConfig)
	default:
		return fmt.Errorf("unsupported output format: %s", f.options.Format)
	}
}

// outputJSON outputs data as formatted JSON
func (f *Formatter) outputJSON(data interface{}) error {
	return f.outputJSONToWriter(os.Stdout, data)
}

// outputJSONToWriter outputs data as formatted JSON to the specified writer
func (f *Formatter) outputJSONToWriter(w io.Writer, data interface{}) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

// outputYAML outputs data as formatted YAML
func (f *Formatter) outputYAML(data interface{}) error {
	return f.outputYAMLToWriter(os.Stdout, data)
}

// outputYAMLToWriter outputs data as formatted YAML to the specified writer
func (f *Formatter) outputYAMLToWriter(w io.Writer, data interface{}) error {
	encoder := yaml.NewEncoder(w)
	encoder.SetIndent(2)
	defer encoder.Close()
	return encoder.Encode(data)
}

// outputTable outputs data as a formatted table
func (f *Formatter) outputTable(data interface{}, config *TableConfig) error {
	return f.outputTableToWriter(os.Stdout, data, config)
}

// outputTableToWriter outputs data as a formatted table to the specified writer
func (f *Formatter) outputTableToWriter(w io.Writer, data interface{}, config *TableConfig) error {
	// Extract items from the data (handle both single items and lists)
	items, err := f.extractItems(data)
	if err != nil {
		return err
	}

	if len(items) == 0 {
		fmt.Fprintln(w, "No items found.")
		return nil
	}

	// Determine which columns to show
	columns := config.Columns
	if f.options.WideOutput {
		columns = append(columns, config.WideCols...)
	} else if f.options.ShowDetails {
		columns = append(columns, config.DetailCols...)
	}

	// Create tabwriter with good formatting
	tw := tabwriter.NewWriter(w, 0, 0, 3, ' ', tabwriter.TabIndent)
	defer tw.Flush()

	// Write headers
	headers := make([]string, len(columns))
	separators := make([]string, len(columns))
	for i, col := range columns {
		headers[i] = col.Header
		separators[i] = strings.Repeat("-", len(col.Header))
	}
	fmt.Fprintf(tw, "%s\n", strings.Join(headers, "\t"))
	fmt.Fprintf(tw, "%s\n", strings.Join(separators, "\t"))

	// Write data rows
	for _, item := range items {
		row := make([]string, len(columns))
		for i, col := range columns {
			value := f.getFieldValue(item, col.Field)
			if value == "" && col.Default != "" {
				value = col.Default
			}

			// Truncate long values if needed
			if col.Width > 0 && len(value) > col.Width {
				value = value[:col.Width-3] + "..."
			}

			row[i] = value
		}
		fmt.Fprintf(tw, "%s\n", strings.Join(row, "\t"))
	}

	return nil
}

// extractItems extracts a slice of items from the data
// Handles both single items and list structures (like RegionList)
func (f *Formatter) extractItems(data interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(data)

	// Handle pointers
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil, fmt.Errorf("nil data provided")
		}
		v = v.Elem()
	}

	// If it's a struct, check for an "Items" field (Kubernetes-style lists)
	if v.Kind() == reflect.Struct {
		itemsField := v.FieldByName("Items")
		if itemsField.IsValid() && itemsField.Kind() == reflect.Slice {
			var items []interface{}
			for i := 0; i < itemsField.Len(); i++ {
				items = append(items, itemsField.Index(i).Interface())
			}
			return items, nil
		}

		// Single item
		return []interface{}{data}, nil
	}

	// If it's already a slice
	if v.Kind() == reflect.Slice {
		var items []interface{}
		for i := 0; i < v.Len(); i++ {
			items = append(items, v.Index(i).Interface())
		}
		return items, nil
	}

	return nil, fmt.Errorf("unsupported data type for table output: %T", data)
}

// getFieldValue extracts a field value from an item using a dot-notation path
// e.g., "metadata.name" or "spec.provider.providerType"
func (f *Formatter) getFieldValue(item interface{}, fieldPath string) string {
	if fieldPath == "" {
		return ""
	}

	parts := strings.Split(fieldPath, ".")
	v := reflect.ValueOf(item)

	// Handle pointers
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return ""
		}
		v = v.Elem()
	}

	for _, part := range parts {
		if v.Kind() != reflect.Struct {
			return ""
		}

		// Handle both direct field names and JSON tag names
		field := f.findField(v, part)
		if !field.IsValid() {
			return ""
		}

		// Handle pointers in the field chain
		if field.Kind() == reflect.Ptr {
			if field.IsNil() {
				return ""
			}
			field = field.Elem()
		}

		v = field
	}

	// Convert the final value to string
	if !v.IsValid() {
		return ""
	}

	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", v.Int())
	case reflect.Bool:
		return fmt.Sprintf("%t", v.Bool())
	default:
		return fmt.Sprintf("%v", v.Interface())
	}
}

// findField finds a struct field by name or JSON tag
func (f *Formatter) findField(v reflect.Value, fieldName string) reflect.Value {
	if v.Kind() != reflect.Struct {
		return reflect.Value{}
	}

	t := v.Type()

	// First try exact field name match
	if field := v.FieldByName(fieldName); field.IsValid() {
		return field
	}

	// Then try case-insensitive match
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		if strings.EqualFold(field.Name, fieldName) {
			return v.Field(i)
		}
	}

	// Finally try JSON tag match
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" {
			// Handle "field,omitempty" format
			tagName := strings.Split(jsonTag, ",")[0]
			if tagName == fieldName {
				return v.Field(i)
			}
		}
	}

	return reflect.Value{}
}
