package preview

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf16"

	"github.com/hostinger/api-cli/deploy/hosting"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/hcdn-connector/pkg/connector"
)

// project is the folder a remote deploy ships, as `hostinger up` would from
// the working directory.
type project struct {
	dir  string        // the link's folder, else the working directory
	link *site.Project // nil: not linked yet
	typ  hosting.Type  // static or nodejs
}

// findProject returns the project around dir, or nil when a remote deploy
// cannot trust dir to be one. A preview knows a port, not a folder, so only
// a folder linked to a website, or an unlinked one that looks like a Node.js
// app, qualifies: any other folder, even the home directory, would deploy as
// static files. A linked WordPress site is not either: an import replaces the
// database.
func findProject(dir string) *project {
	link, linkDir, err := site.Find(dir)
	if err != nil {
		return nil
	}
	if link == nil {
		if typ, _, err := hosting.Detect(dir); err == nil && typ == hosting.NodeJS {
			return &project{dir: dir, typ: typ}
		}
		return nil
	}
	if link.Kind != "" && link.Kind != "website" {
		return nil
	}
	typ, err := hosting.LinkedType(link, filepath.Join(linkDir, link.Path))
	if err != nil || (typ != hosting.Static && typ != hosting.NodeJS) {
		return nil
	}
	return &project{dir: linkDir, link: link, typ: typ}
}

// currentProject is the project of the working directory, which `hostinger
// up` deploys from.
func currentProject() *project {
	wd, err := hosting.Getwd()
	if err != nil {
		return nil
	}
	return findProject(wd)
}

// linkedElsewhere reports a link to another website than username's domain.
func (p *project) linkedElsewhere(username, domain string) bool {
	return p.link != nil && (!strings.EqualFold(p.link.Domain, domain) || (p.link.Username != "" && p.link.Username != username))
}

// deployOptions are `hostinger up`'s options for deploying p to domain: a
// linked project deploys from its link, and a first deploy goes to domain
// and links it. overwrite allows replacing a website's content.
func (p *project) deployOptions(domain string, overwrite bool, message string) hosting.Options {
	opts := hosting.Options{Path: ".", Remote: true, Yes: overwrite, Message: message}
	if p.link == nil {
		opts.Domain = domain
	}
	return opts
}

// target is the website the project deploys to, or "" before its first deploy.
func (p *project) target() string {
	if p.link == nil {
		return ""
	}
	return p.link.Domain
}

// appMeta is what hPanel shows of the project before a deploy: the linked
// website, the type and the folder's name (never its path). It is JSON in
// printable ASCII within the connector's cap: an unusual folder name is
// escaped, and shortened when it would not fit.
func (p *project) appMeta() string {
	m := struct {
		Target string       `json:"target,omitempty"`
		Type   hosting.Type `json:"type"`
		Folder string       `json:"folder,omitempty"`
	}{Target: p.target(), Type: p.typ, Folder: folderName(filepath.Base(p.dir), maxFolderBytes)}
	for {
		raw, _ := json.Marshal(m) // strings only: cannot fail
		// Without a folder it fits: a link's domain is a hostname.
		if out := asciiJSON(raw); len(out) <= connector.MaxAppMetaBytes || m.Folder == "" {
			return out
		}
		m.Folder = folderName(m.Folder, len(m.Folder)-1)
	}
}

// maxFolderBytes is the longest folder name hPanel is told.
const maxFolderBytes = 255

// folderName is name, cut to at most limit bytes, as hPanel can show it, or
// "" for none. The edge's API drops the whole metadata, target and type with
// it, for a name that is longer than maxFolderBytes, starts or ends with a
// space, or holds a character that does not print or is special in HTML or
// a path. Spaces become plain ones and other characters that do not print
// are left out, which changes nothing on screen; a name with a special
// character is not shown at all, rather than shown wrong.
func folderName(name string, limit int) string {
	name = strings.Map(func(r rune) rune {
		switch {
		case unicode.IsSpace(r):
			return ' '
		case !unicode.IsPrint(r):
			return -1
		}
		return r
	}, name)
	if len(name) > limit {
		name = strings.ToValidUTF8(name[:limit], "") // drops a character cut in two
	}
	name = strings.TrimSpace(name)
	if name == "." || name == ".." || strings.ContainsAny(name, `/\<>"'&`+"`") {
		return ""
	}
	return name
}

// asciiJSON escapes every byte of encoded JSON outside printable ASCII as
// \uXXXX, which leaves the value it decodes to unchanged. encoding/json has
// already escaped control characters and replaced invalid UTF-8.
func asciiJSON(raw []byte) string {
	var b strings.Builder
	for _, r := range string(raw) {
		if r >= 0x20 && r < 0x7f {
			b.WriteRune(r)
			continue
		}
		if r > 0xffff {
			hi, lo := utf16.EncodeRune(r)
			fmt.Fprintf(&b, `\u%04x\u%04x`, hi, lo)
			continue
		}
		fmt.Fprintf(&b, `\u%04x`, r)
	}
	return b.String()
}
