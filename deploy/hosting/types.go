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
