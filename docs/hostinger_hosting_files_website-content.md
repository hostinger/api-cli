## hostinger hosting files website-content

Get website file content

### Synopsis

Get a single file's content, relative to a website's document root.

Read-only; refuses symlinks, oversized files, non-text file types, and files identified as
containing secrets (e.g. credential files) — none of these are returned by this endpoint.

```
hostinger hosting files website-content <username> <domain> [flags]
```

### Examples

```
  hostinger hosting files website-content u123456789 mydomain.tld --path index.php
```

### Options

```
      --from-line int   Line offset to start reading from.
  -h, --help            help for website-content
      --max-lines int   Max number of lines to return. (default 5000)
      --path string     File path, relative to the document root.
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

