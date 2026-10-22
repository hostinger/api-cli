## hostinger agency-hosting databases list-website

List website databases

### Synopsis

Returns a paginated list of MySQL databases created for an Agency Plan website.

Each entry includes the database's non-system users.

```
hostinger agency-hosting databases list-website <website_uid> [flags]
```

### Examples

```
  hostinger agency-hosting databases list-website zpwlGlp19
```

### Options

```
      --all            Fetch every page and print all items
  -h, --help           help for list-website
      --page int       Page number
      --per-page int   Number of items per page (default 25)
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

* [hostinger agency-hosting databases](hostinger_agency-hosting_databases.md)	 - Databases commands

