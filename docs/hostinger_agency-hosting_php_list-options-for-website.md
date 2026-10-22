## hostinger agency-hosting php list-options-for-website

List PHP options for a website

### Synopsis

Lists the php.ini directives that can be configured for an Agency Plan website, each with its default, the value currently in effect, and the values it accepts.

```
hostinger agency-hosting php list-options-for-website <website_uid> [flags]
```

### Examples

```
  hostinger agency-hosting php list-options-for-website zpwlGlp19
```

### Options

```
  -h, --help   help for list-options-for-website
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

