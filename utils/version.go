package utils

import (
	"bytes"
	"fmt"
	"strconv"
)

type Version struct {
	Major, Minor, Patch int
	Build               string
}

var (
	Build string
	Major string
	Minor string
	Patch string

	CLIVersion Version
)

func init() {
	if Build != "" {
		CLIVersion.Build = Build
	}
	if Major != "" {
		i, _ := strconv.Atoi(Major)
		CLIVersion.Major = i
	}
	if Minor != "" {
		i, _ := strconv.Atoi(Minor)
		CLIVersion.Minor = i
	}
	if Patch != "" {
		i, _ := strconv.Atoi(Patch)
		CLIVersion.Patch = i
	}
}

func (v Version) String(withBuild bool) string {
	var buffer bytes.Buffer
	if withBuild {
		buffer.WriteString(fmt.Sprintf("%d.%d.%d (Build: %s)", v.Major, v.Minor, v.Patch, v.Build))
	} else {
		buffer.WriteString(fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch))
	}

	return buffer.String()
}
