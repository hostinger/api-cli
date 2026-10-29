package output

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/aquasecurity/table"
	"github.com/liamg/tml"
)

func printTable(body []byte) {
	err := displayJSONAsTables(body)
	if err != nil {
		fmt.Printf(tml.Sprintf("<red>Error displaying JSON: %v</red>\n"), err)
		return
	}
}

func displayJSONAsTables(body []byte) error {
	var data interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	data = safeJSON(data)

	switch v := data.(type) {
	case []interface{}:
		// Array - each item in separate table row
		return displayArrayAsTable(v)
	case map[string]interface{}:
		// Object - handle special cases like paginated responses
		return displayObjectAsSeparateTables(v)
	default:
		return fmt.Errorf("unsupported JSON structure: expected object or array of objects")
	}
}

func displayObjectAsSeparateTables(obj map[string]interface{}) error {
	// Check if this is a paginated response with 'data' and 'meta' fields
	if data, hasData := obj["data"]; hasData {
		if meta, hasMeta := obj["meta"]; hasMeta {
			// Display data first
			printSectionHeader("data")
			if dataArray, ok := data.([]interface{}); ok {
				if err := displayArrayAsTable(dataArray); err != nil {
					return err
				}
			} else {
				if err := displayValue(data); err != nil {
					return err
				}
			}

			// Display meta information
			fmt.Print(tml.Sprintf("\n"))
			printSectionHeader("meta")
			if metaObj, ok := meta.(map[string]interface{}); ok {
				return displayMetaAsTable(metaObj)
			} else {
				return displayValue(meta)
			}
		}
	}

	// Get sorted keys for consistent ordering
	keys := make([]string, 0, len(obj))
	for key := range obj {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for i, key := range keys {
		value := obj[key]

		// Print section header
		if i > 0 {
			fmt.Print(tml.Sprintf("\n"))
		}
		printSectionHeader(key)

		if err := displayValue(value); err != nil {
			return err
		}
	}

	return nil
}

func displayValue(value interface{}) error {
	switch v := value.(type) {
	case []interface{}:
		// Array - display as table with rows
		return displayArrayAsTable(v)
	case map[string]interface{}:
		// Nested object - display as key-value table
		return displayObjectAsKeyValueTable(v)
	default:
		// Simple value - display as single value
		fmt.Print(tml.Sprintf("%s\n", formatValue(v)))
		return nil
	}
}

func displayMetaAsTable(meta map[string]interface{}) error {
	t := table.New(os.Stdout)
	t.SetHeaders("Property", "Value")

	// Get sorted keys
	keys := make([]string, 0, len(meta))
	for key := range meta {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		t.AddRow(key, formatValue(meta[key]))
	}

	t.Render()
	return nil
}

func displayObjectAsKeyValueTable(obj map[string]interface{}) error {
	return displayAsTree(obj, "", true)
}

func displayAsTree(data interface{}, prefix string, isRoot bool) error {
	switch v := data.(type) {
	case map[string]interface{}:
		return displayObjectAsTree(v, prefix, isRoot)
	case []interface{}:
		return displayArrayAsTree(v, prefix, isRoot)
	default:
		fmt.Printf("%s%s\n", prefix, formatValue(v))
		return nil
	}
}

func displayObjectAsTree(obj map[string]interface{}, prefix string, isRoot bool) error {
	if len(obj) == 0 {
		fmt.Printf("%s%s\n", prefix, tml.Sprintf("<italic>(empty object)</italic>"))
		return nil
	}

	// Get sorted keys
	keys := make([]string, 0, len(obj))
	for key := range obj {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for i, key := range keys {
		isLast := i == len(keys)-1
		value := obj[key]

		// Determine tree characters
		var keyPrefix, childPrefix string
		if isRoot {
			keyPrefix = ""
			childPrefix = ""
		} else {
			if isLast {
				keyPrefix = prefix + "└── "
				childPrefix = prefix + "    "
			} else {
				keyPrefix = prefix + "├── "
				childPrefix = prefix + "│   "
			}
		}

		// Display the key
		if isRoot {
			fmt.Printf("%s:\n", tml.Sprintf("<bold>%s</bold>", key))
		} else {
			fmt.Printf("%s%s: ", keyPrefix, tml.Sprintf("<bold>%s</bold>", key))
		}

		// Handle the value
		switch val := value.(type) {
		case map[string]interface{}:
			if !isRoot {
				fmt.Print("\n")
			}
			if err := displayAsTree(val, childPrefix, false); err != nil {
				return err
			}
		case []interface{}:
			if !isRoot {
				fmt.Print("\n")
			}
			if err := displayAsTree(val, childPrefix, false); err != nil {
				return err
			}
		default:
			if isRoot {
				fmt.Printf("  %s\n", formatValue(val))
			} else {
				fmt.Printf("%s\n", formatValue(val))
			}
		}
	}

	return nil
}

func displayArrayAsTree(arr []interface{}, prefix string, isRoot bool) error {
	if len(arr) == 0 {
		fmt.Printf("%s%s\n", prefix, tml.Sprintf("<italic>(empty array)</italic>"))
		return nil
	}

	for i, item := range arr {
		isLast := i == len(arr)-1

		// Determine tree characters
		var itemPrefix, childPrefix string
		if isLast {
			itemPrefix = prefix + "└── "
			childPrefix = prefix + "    "
		} else {
			itemPrefix = prefix + "├── "
			childPrefix = prefix + "│   "
		}

		// Display array index
		fmt.Printf("%s[%d]: ", itemPrefix, i)

		// Handle the item
		switch val := item.(type) {
		case map[string]interface{}:
			fmt.Print("\n")
			if err := displayAsTree(val, childPrefix, false); err != nil {
				return err
			}
		case []interface{}:
			fmt.Print("\n")
			if err := displayAsTree(val, childPrefix, false); err != nil {
				return err
			}
		default:
			fmt.Printf("%s\n", formatValue(val))
		}
	}

	return nil
}

func displayArrayAsTable(arr []interface{}) error {
	if len(arr) == 0 {
		fmt.Print(tml.Sprintf("<italic>(empty array)</italic>\n"))
		return nil
	}

	// Check if array contains objects or simple values
	if isArrayOfObjects(arr) {
		return displayObjectArrayAsTable(arr)
	} else {
		return displaySimpleArrayAsTable(arr)
	}
}

func isArrayOfObjects(arr []interface{}) bool {
	for _, item := range arr {
		if _, ok := item.(map[string]interface{}); ok {
			return true
		}
	}
	return false
}

func displayObjectArrayAsTable(arr []interface{}) error {
	// Collect all unique keys from all objects
	allKeys := make(map[string]bool)
	for _, item := range arr {
		if objMap, ok := item.(map[string]interface{}); ok {
			for key := range objMap {
				allKeys[key] = true
			}
		}
	}

	// Convert to sorted slice
	keys := make([]string, 0, len(allKeys))
	for key := range allKeys {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Create table
	t := table.New(os.Stdout)
	t.SetHeaders(keys...)

	// Add each object as a row
	for _, item := range arr {
		objMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		row := make([]string, len(keys))
		for i, key := range keys {
			if value, exists := objMap[key]; exists {
				row[i] = formatCellValue(value)
			} else {
				row[i] = "-"
			}
		}
		t.AddRow(row...)
	}

	t.Render()
	return nil
}

func displaySimpleArrayAsTable(arr []interface{}) error {
	// Simple array - display as single column table
	t := table.New(os.Stdout)
	t.SetHeaders("Value")

	for _, item := range arr {
		t.AddRow(formatValue(item))
	}

	t.Render()
	return nil
}

func printSectionHeader(text string) {
	fmt.Print(tml.Sprintf("<bold><blue>=== %s ===</blue></bold>\n", strings.ToUpper(text)))
}

// formatCellValue formats values for table cells (more compact)
func formatCellValue(value interface{}) string {
	if value == nil {
		return tml.Sprintf("<italic>null</italic>")
	}

	switch v := value.(type) {
	case string:
		if v == "" {
			return tml.Sprintf("<italic>(empty)</italic>")
		}
		return v
	case bool:
		return tml.Sprintf("<magenta>%s</magenta>", strconv.FormatBool(v))
	case float64:
		// Check if it's actually an integer
		if v == float64(int64(v)) {
			return tml.Sprintf("<cyan>%s</cyan>", strconv.FormatInt(int64(v), 10))
		}
		return tml.Sprintf("<cyan>%s</cyan>", strconv.FormatFloat(v, 'f', -1, 64))
	case int:
		return tml.Sprintf("<cyan>%s</cyan>", strconv.Itoa(v))
	case int64:
		return tml.Sprintf("<cyan>%s</cyan>", strconv.FormatInt(v, 10))
	case []interface{}, map[string]interface{}:
		// For nested structures in table cells, format as tree string
		return formatNestedAsTreeString(v, "")
	default:
		return fmt.Sprintf("%v", v)
	}
}

// formatNestedAsTreeString formats nested structures as a compact tree string for table cells
func formatNestedAsTreeString(data interface{}, prefix string) string {
	var result strings.Builder

	switch v := data.(type) {
	case map[string]interface{}:
		if len(v) == 0 {
			return "{}"
		}

		// Get sorted keys
		keys := make([]string, 0, len(v))
		for key := range v {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for i, key := range keys {
			isLast := i == len(keys)-1
			var keyPrefix, childPrefix string

			if isLast {
				keyPrefix = prefix + "└── "
				childPrefix = prefix + "    "
			} else {
				keyPrefix = prefix + "├── "
				childPrefix = prefix + "│   "
			}

			if i > 0 {
				result.WriteString("\n")
			}

			value := v[key]
			switch val := value.(type) {
			case map[string]interface{}, []interface{}:
				result.WriteString(fmt.Sprintf("%s%s:", keyPrefix, key))
				result.WriteString("\n")
				result.WriteString(formatNestedAsTreeString(val, childPrefix))
			default:
				// Format the whole node (key + value)
				nodeText := fmt.Sprintf("%s%s: %s", keyPrefix, key, formatValuePlain(val))
				result.WriteString(truncateNodeIfNeeded(nodeText, keyPrefix, key, formatValuePlain(val)))
			}
		}

	case []interface{}:
		if len(v) == 0 {
			return "[]"
		}

		for i, item := range v {
			isLast := i == len(v)-1
			var itemPrefix, childPrefix string

			if isLast {
				itemPrefix = prefix + "└── "
				childPrefix = prefix + "    "
			} else {
				itemPrefix = prefix + "├── "
				childPrefix = prefix + "│   "
			}

			if i > 0 {
				result.WriteString("\n")
			}

			switch val := item.(type) {
			case map[string]interface{}, []interface{}:
				result.WriteString(fmt.Sprintf("%s[%d]:", itemPrefix, i))
				result.WriteString("\n")
				result.WriteString(formatNestedAsTreeString(val, childPrefix))
			default:
				// Format the whole node (index + value)
				nodeText := fmt.Sprintf("%s[%d]: %s", itemPrefix, i, formatValuePlain(val))
				result.WriteString(truncateNodeIfNeeded(nodeText, itemPrefix, fmt.Sprintf("[%d]", i), formatValuePlain(val)))
			}
		}
	}

	return result.String()
}

// truncateNodeIfNeeded truncates the whole node if it's longer than 32 chars and the value contains spaces
func truncateNodeIfNeeded(fullNode, prefix, label, value string) string {
	// Check if the value (not the whole node) contains spaces
	if len(fullNode) > 32 && strings.Contains(value, " ") {
		return fullNode[:32] + "..."
	}
	return fullNode
}

// formatValuePlain formats values without color codes for tree strings
func formatValuePlain(value interface{}) string {
	if value == nil {
		return "null"
	}

	switch v := value.(type) {
	case string:
		if v == "" {
			return "(empty)"
		}
		return v
	case bool:
		return strconv.FormatBool(v)
	case float64:
		if v == float64(int64(v)) {
			return strconv.FormatInt(int64(v), 10)
		}
		return strconv.FormatFloat(v, 'f', -1, 64)
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// formatValue formats values for detailed display
func formatValue(value interface{}) string {
	if value == nil {
		return tml.Sprintf("<italic>null</italic>")
	}

	switch v := value.(type) {
	case string:
		if v == "" {
			return tml.Sprintf("<italic>(empty)</italic>")
		}
		return v
	case bool:
		return tml.Sprintf("<magenta>%s</magenta>", strconv.FormatBool(v))
	case float64:
		// Check if it's actually an integer
		if v == float64(int64(v)) {
			return tml.Sprintf("<cyan>%s</cyan>", strconv.FormatInt(int64(v), 10))
		}
		return tml.Sprintf("<cyan>%s</cyan>", strconv.FormatFloat(v, 'f', -1, 64))
	case int:
		return tml.Sprintf("<cyan>%s</cyan>", strconv.Itoa(v))
	case int64:
		return tml.Sprintf("<cyan>%s</cyan>", strconv.FormatInt(v, 10))
	case []interface{}:
		if len(v) == 0 {
			return "[]"
		}
		// For arrays, show a summary
		return fmt.Sprintf("Array with %d items", len(v))
	case map[string]interface{}:
		if len(v) == 0 {
			return "{}"
		}
		// For objects, show a summary
		return fmt.Sprintf("Object with %d properties", len(v))
	default:
		return fmt.Sprintf("%v", v)
	}
}
