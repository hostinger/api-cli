## hostinger agency-hosting php list-extensions-for-website

List PHP extensions for a website

### Synopsis

Lists every PHP extension available to an Agency Plan website and whether it is currently enabled.

```
hostinger agency-hosting php list-extensions-for-website <website_uid> [flags]
```

### Examples

```
  hostinger agency-hosting php list-extensions-for-website zpwlGlp19
```

### Options

```
  -h, --help   help for list-extensions-for-website
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

* [hostinger agency-hosting php](hostinger_agency-hosting_php.md)	 - PHP commands

