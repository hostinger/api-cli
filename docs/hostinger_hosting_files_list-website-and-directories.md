## hostinger hosting files list-website-and-directories

List website files and directories

### Synopsis

List files and directories under a website's document root.

Use `directory` to browse a subdirectory relative to the document root. Symlinked entries
are listed but never traversed into or resolved.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting files list-website-and-directories [domain] [flags]
```

### Examples

```
  hostinger hosting files list-website-and-directories mydomain.tld
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
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --domain string     Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger hosting files](hostinger_hosting_files.md)	 - Files commands

