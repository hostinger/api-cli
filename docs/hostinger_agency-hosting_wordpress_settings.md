## hostinger agency-hosting wordpress settings

Get WordPress settings

### Synopsis

Returns the current WordPress settings for an Agency Plan website: installed core version,
LiteSpeed Cache plugin status, object cache status, and maintenance mode status.

```
hostinger agency-hosting wordpress settings <website_uid> [flags]
```

### Examples

```
  hostinger agency-hosting wordpress settings zpwlGlp19
```

### Options

```
  -h, --help   help for settings
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

* [hostinger agency-hosting wordpress](hostinger_agency-hosting_wordpress.md)	 - WordPress commands

