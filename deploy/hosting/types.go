// Package hosting deploys a local directory to a Hostinger website.
package hosting

import "fmt"

type Type string

const (
	Static          Type = "static" // HTML, CSS, JavaScript and PHP served as files
	NodeJS          Type = "nodejs" // built and run by the Node.js runtime
	WordPressSite   Type = "wordpress-site"
	WordPressPlugin Type = "wordpress-plugin"
	WordPressTheme  Type = "wordpress-theme"
)

var DeployableTypes = []Type{Static, NodeJS}

func ParseType(s string) (Type, error) {
	for _, t := range DeployableTypes {
		if string(t) == s {
			return t, nil
		}
	}
	return "", fmt.Errorf("unknown type %q (one of: static, nodejs)", s)
}

// HeadlessKinds are the headless websites up creates: empty installs nothing,
// wordpress installs WordPress in /cms, ecommerce and pocketbase are hosts for
// a Hostinger Ecommerce store or a PocketBase backend.
var HeadlessKinds = []string{"empty", "wordpress", "ecommerce", "pocketbase"}

func ParseHeadless(s string) (string, error) {
	for _, k := range HeadlessKinds {
		if k == s {
			return k, nil
		}
	}
	return "", fmt.Errorf("unknown headless website %q (one of: empty, wordpress, ecommerce, pocketbase)", s)
}
