package site

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Context is what `hostinger link --show` prints: each value and its source.
type Context struct {
	Domain         string   `json:"domain,omitempty"`
	DomainSource   string   `json:"domainSource"`
	Username       string   `json:"username,omitempty"`
	UsernameSource string   `json:"usernameSource,omitempty"`
	OrderID        string   `json:"orderId,omitempty"`
	OrderSource    string   `json:"orderSource,omitempty"`
	LinkFile       string   `json:"linkFile,omitempty"`
	Link           *Project `json:"link,omitempty"`
	ConfigFile     string   `json:"configFile,omitempty"`
}

// Explain resolves like the generated commands, without prompting.
func Explain(cmd *cobra.Command) (*Context, error) {
	c := &Context{DomainSource: "not set", ConfigFile: viper.ConfigFileUsed()}
	p, dir, err := Find(".")
	if err != nil {
		return nil, err
	}
	if p != nil {
		c.Link = p
		c.LinkFile = Path(dir)
	}

	flag := func(name string) string {
		if cmd == nil {
			return ""
		}
		if f := cmd.Flags().Lookup(name); f != nil && f.Changed {
			return f.Value.String()
		}
		return ""
	}
	switch {
	case flag("domain") != "":
		c.Domain, c.DomainSource = NormalizeDomain(flag("domain")), "--domain"
	case os.Getenv("HOSTINGER_DOMAIN") != "":
		c.Domain, c.DomainSource = NormalizeDomain(os.Getenv("HOSTINGER_DOMAIN")), "HOSTINGER_DOMAIN"
	case p != nil:
		c.Domain, c.DomainSource = p.Domain, "link"
	case viper.GetString("default_domain") != "":
		c.Domain, c.DomainSource = NormalizeDomain(viper.GetString("default_domain")), "default_domain in "+c.ConfigFile+"; not used by `hostinger up`"
	}

	switch {
	case flag("username") != "":
		c.Username, c.UsernameSource = flag("username"), "--username"
	case os.Getenv("HOSTINGER_USERNAME") != "":
		c.Username, c.UsernameSource = os.Getenv("HOSTINGER_USERNAME"), "HOSTINGER_USERNAME"
	case p != nil && p.Username != "" && strings.EqualFold(p.Domain, c.Domain):
		c.Username, c.UsernameSource = p.Username, "link"
	case c.Domain != "":
		c.UsernameSource = "looked up from " + c.Domain + " when needed"
	}

	switch {
	case flag("order-id") != "":
		c.OrderID, c.OrderSource = flag("order-id"), "--order-id"
	case os.Getenv("HOSTINGER_ORDER_ID") != "":
		c.OrderID, c.OrderSource = os.Getenv("HOSTINGER_ORDER_ID"), "HOSTINGER_ORDER_ID"
	case p != nil && p.OrderID != 0:
		c.OrderID, c.OrderSource = strconv.Itoa(p.OrderID), "link"
	}
	return c, nil
}

func (c *Context) Lines() []string {
	row := func(label, value, from string) string {
		if value == "" {
			value = "-"
		}
		if from != "" {
			return fmt.Sprintf("%-9s %s  (%s)", label, value, from)
		}
		return fmt.Sprintf("%-9s %s", label, value)
	}
	out := []string{
		row("Website", c.Domain, c.DomainSource),
		row("Account", c.Username, c.UsernameSource),
		row("Order", c.OrderID, c.OrderSource),
	}
	if c.Link != nil {
		folder := c.Link.Path
		if folder == "" {
			folder = "."
		}
		typ := c.Link.Type
		if c.Link.Headless {
			typ += " (headless website)"
		}
		out = append(out,
			row("Type", typ, fromLink(c.Link.Type)),
			row("Source", c.Link.Source, fromLink(c.Link.Source)),
			row("Folder", folder, ""),
		)
		if n := c.Link.Node; n != nil {
			build := n.AppType
			if n.NodeVersion != 0 {
				build += fmt.Sprintf(" on Node.js %d", n.NodeVersion)
			}
			if n.EntryFile != "" {
				build += ", entry " + n.EntryFile
			}
			if n.BuildScript != "" {
				build += ", build script " + n.BuildScript
			}
			out = append(out, row("Build", build, "link"))
		}
		if b := c.Link.Build; b != nil && b.UUID != "" {
			state := b.State
			if state == "" {
				state = "started"
			}
			out = append(out, row("Build id", fmt.Sprintf("%s  (%s, %s)", b.UUID, state, b.At.Local().Format("2006-01-02 15:04")), "link"))
		}
		out = append(out, row("Link", c.LinkFile, ""))
	} else {
		out = append(out, row("Link", "none (run `hostinger up` or `hostinger link`)", ""))
	}
	return out
}

func fromLink(value string) string {
	if value == "" {
		return ""
	}
	return "link"
}
