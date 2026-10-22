## hostinger hosting databases list

List account databases

### Synopsis

Returns a paginated list of databases for the specified account.

Use the domain and is_assigned filters to find databases assigned to a specific domain.

```
hostinger hosting databases list <username> [flags]
```

### Examples

```
  hostinger hosting databases list u123456789
```

### Options

```
      --all             Fetch every page and print all items
      --domain string   Filter by domain name (case-insensitive substring match)
  -h, --help            help for list
      --is-assigned     When used with domain, return only databases assigned to that domain.
      --page int        Page number
      --per-page int    Number of items per page (default 25)
      --search string   Search databases by name, user, or creation date.
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

* [hostinger hosting databases](hostinger_hosting_databases.md)	 - Databases commands

