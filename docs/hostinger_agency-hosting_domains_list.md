## hostinger agency-hosting domains list

List domains

### Synopsis

Returns a paginated list of domains associated with Agency Plan websites accessible to the authenticated client.

Use the website_uuids filter to narrow results to specific websites.

```
hostinger agency-hosting domains list [flags]
```

### Examples

```
  hostinger agency-hosting domains list
```

### Options

```
      --all                     Fetch every page and print all items
  -h, --help                    help for list
      --page int                Page number
      --per-page int            Number of items per page (default 25)
      --website-uuids strings   Filter by website UIDs
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

* [hostinger agency-hosting domains](hostinger_agency-hosting_domains.md)	 - Domains commands

