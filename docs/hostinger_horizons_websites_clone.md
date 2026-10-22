## hostinger horizons websites clone

Clone website

### Synopsis

Clone a Hostinger Horizons website into a new website.\n
Use this tool when the user wants a copy of an existing website, for example to try out
changes without touching the original.\n
This tool returns the ID and URL of the newly created copy.
The original website is left untouched.\n
To edit the copy, use the `Edit website` tool with the returned website ID, or the user can
open the provided website URL in Hostinger Horizons interface.

```
hostinger horizons websites clone <website-id> [flags]
```

### Examples

```
  hostinger horizons websites clone 123e4567-e89b-12d3-a456-426614174000
```

### Options

```
  -h, --help   help for clone
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

* [hostinger horizons websites](hostinger_horizons_websites.md)	 - Websites commands

