package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/hostinger/api-cli/utils"
	"github.com/itchyny/gojq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.yaml.in/yaml/v3"
	"golang.org/x/term"
)

// Formats are the values --format accepts.
var Formats = []string{"table", "json", "yaml", "tree"}

// stdout (nil: os.Stdout at write time) and its terminal check are package
// vars for tests.
var (
	stdout           io.Writer
	stdoutIsTerminal = func() bool { return term.IsTerminal(int(os.Stdout.Fd())) }
)

func out() io.Writer {
	if stdout != nil {
		return stdout
	}
	return os.Stdout
}

// Format prints a successful response in the selected format, or returns an
// *utils.APIError for an error status. --jq filters the JSON first.
func Format(cmd *cobra.Command, body []byte, statusCode int) error {
	if statusCode >= http.StatusBadRequest {
		return utils.NewAPIError(statusCode, body)
	}
	format, err := SelectedFormat(cmd)
	if err != nil {
		return err
	}
	if len(bytes.TrimSpace(body)) == 0 {
		return nil // e.g. 204 from a delete: nothing to print or filter
	}
	if expr, _ := cmd.Flags().GetString("jq"); expr != "" {
		return printJQ(expr, body)
	}

	switch format {
	case "json":
		fmt.Fprintln(out(), string(body))
	case "yaml":
		return printYAML(body)
	case "tree":
		printTreeEnhanced(body)
	default:
		printTable(body)
	}
	return nil
}

// SelectedFormat resolves the output format: --format, then the "format"
// config key / environment variable, then table on a terminal and JSON when
// the output is piped or redirected.
func SelectedFormat(cmd *cobra.Command) (string, error) {
	format, _ := cmd.Flags().GetString("format")
	if format == "" {
		format = viper.GetString("format")
	}
	if format == "" {
		if stdoutIsTerminal() {
			return "table", nil
		}
		return "json", nil
	}
	for _, f := range Formats {
		if f == format {
			return format, nil
		}
	}
	return "", utils.FlagErrorf("invalid format %q (one of: table, json, yaml, tree)", format)
}

// printJQ applies a jq expression to the response. Strings print raw (like
// jq -r); every other result prints as compact JSON, one per line. On a
// terminal raw strings lose control characters; piped, they stay byte-exact.
func printJQ(expr string, body []byte) error {
	q, err := gojq.Parse(expr)
	if err != nil {
		return utils.FlagErrorf("invalid --jq expression: %v", err)
	}
	var input any
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.UseNumber()
	if err := dec.Decode(&input); err != nil {
		return fmt.Errorf("--jq: response is not JSON: %w", err)
	}
	iter := q.Run(normalizeNumbers(input))
	enc := json.NewEncoder(out())
	enc.SetEscapeHTML(false) // print & < > like jq does
	terminal := stdoutIsTerminal()
	for {
		v, ok := iter.Next()
		if !ok {
			return nil
		}
		if err, ok := v.(error); ok {
			return fmt.Errorf("--jq: %w", err)
		}
		if s, ok := v.(string); ok {
			if terminal {
				s = utils.SafeText(s)
			}
			fmt.Fprintln(out(), s)
			continue
		}
		if err := enc.Encode(v); err != nil {
			return err
		}
	}
}

// normalizeNumbers converts json.Number into the int/float64 values gojq
// understands, keeping integers exact.
func normalizeNumbers(v any) any {
	switch t := v.(type) {
	case json.Number:
		if i, err := t.Int64(); err == nil {
			return int(i)
		}
		f, _ := t.Float64()
		return f
	case map[string]any:
		for k, x := range t {
			t[k] = normalizeNumbers(x)
		}
	case []any:
		for i, x := range t {
			t[i] = normalizeNumbers(x)
		}
	}
	return v
}

// safeJSON strips terminal control characters from every key and string of a
// decoded response before table/tree rendering, so text written by other
// people (form submissions, orders) cannot drive the terminal.
func safeJSON(v any) any {
	switch t := v.(type) {
	case string:
		return utils.SafeText(t)
	case map[string]any:
		m := make(map[string]any, len(t))
		for k, x := range t {
			m[utils.SafeText(k)] = safeJSON(x)
		}
		return m
	case []any:
		for i, x := range t {
			t[i] = safeJSON(x)
		}
	}
	return v
}

// printYAML re-encodes the JSON response as block-style YAML, keeping the
// API's key order (JSON is valid YAML, so it decodes straight into a node).
func printYAML(body []byte) error {
	var node yaml.Node
	if err := yaml.Unmarshal(body, &node); err != nil {
		return fmt.Errorf("converting response to YAML: %w", err)
	}
	blockStyle(&node)
	enc := yaml.NewEncoder(out())
	enc.SetIndent(2)
	if err := enc.Encode(&node); err != nil {
		return err
	}
	return enc.Close()
}

func blockStyle(n *yaml.Node) {
	n.Style &^= yaml.FlowStyle
	if n.Kind == yaml.ScalarNode && n.Style&(yaml.DoubleQuotedStyle|yaml.SingleQuotedStyle) != 0 && n.Tag == "!!str" {
		n.Style = 0 // let the encoder quote only where YAML needs it
	}
	for _, c := range n.Content {
		blockStyle(c)
	}
}
