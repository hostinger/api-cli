package output

import (
	"encoding/json"
	"fmt"
	"github.com/pterm/pterm"
	"sort"
)

func printTreeEnhanced(body []byte) {
	var data interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		pterm.Error.Printf("Error parsing JSON: %v\n", err)
		return
	}

	tree := createTreeNodeEnhanced("JSON", safeJSON(data))
	pterm.DefaultTree.WithRoot(tree).Render()
}

func createTreeNodeEnhanced(name string, value interface{}) pterm.TreeNode {
	node := pterm.TreeNode{}

	switch v := value.(type) {
	case map[string]interface{}:
		// Object: create child nodes for each key-value pair
		node.Text = pterm.Sprintf("%s",
			pterm.Blue(name),
		)

		// Sort keys for consistent order
		var keys []string
		for key := range v {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, key := range keys {
			childNode := createTreeNodeEnhanced(key, v[key])
			node.Children = append(node.Children, childNode)
		}

	case []interface{}:
		// Array: create child nodes for each element
		node.Text = pterm.Sprintf("%s %s",
			pterm.Blue(name),
			pterm.Gray(fmt.Sprintf("(%d items)", len(v))))

		for i, item := range v {
			childName := fmt.Sprintf("[%d]", i)
			childNode := createTreeNodeEnhanced(childName, item)
			node.Children = append(node.Children, childNode)
		}

	case string:
		node.Text = pterm.Sprintf("%s: %s",
			pterm.Cyan(name),
			pterm.Green(fmt.Sprintf("\"%s\"", v)))

	case float64:
		var valueStr string
		if v == float64(int64(v)) {
			valueStr = fmt.Sprintf("%.0f", v)
		} else {
			valueStr = fmt.Sprintf("%.2f", v)
		}
		node.Text = pterm.Sprintf("%s: %s",
			pterm.Cyan(name),
			pterm.Yellow(valueStr))

	case bool:
		color := pterm.Red
		if v {
			color = pterm.Green
		}
		node.Text = pterm.Sprintf("%s: %s",
			pterm.Cyan(name),
			color(fmt.Sprintf("%t", v)))

	case nil:
		node.Text = pterm.Sprintf("%s: %s",
			pterm.Cyan(name),
			pterm.Gray("null"))

	default:
		node.Text = pterm.Sprintf("%s: %s",
			pterm.Cyan(name),
			pterm.White(fmt.Sprintf("%v", v)))
	}

	return node
}
