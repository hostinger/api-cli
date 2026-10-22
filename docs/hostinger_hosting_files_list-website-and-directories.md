## hostinger hosting files list-website-and-directories

List website files and directories

### Synopsis

List files and directories under a website's document root.

Use `directory` to browse a subdirectory relative to the document root. Symlinked entries
are listed but never traversed into or resolved.

```
hostinger hosting files list-website-and-directories <username> <domain> [flags]
```

### Examples

```
  hostinger hosting files list-website-and-directories u123456789 mydomain.tld
```

### Options

```
      --directory string     Directory path to check
      --file-types strings   Filter by entry type, e.g. file,directory. Omit for all types. (one of: file, directory, symlink, other)
  -h, --help                 help for list-website-and-directories
      --max-depth int        How many directory levels deep to recurse. (default 5)
      --max-items int        Max number of entries to return in this page. (default 1000)
      --offset int           Number of entries to skip. Page with offset + item count until reaching total_items.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger hosting files](hostinger_hosting_files.md)	 - Files commands

