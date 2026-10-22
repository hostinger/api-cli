## hostinger vps data-centers list

Get data center list

### Synopsis

Retrieve all available data centers.

Use this endpoint to view location options before deploying VPS instances.

```
hostinger vps data-centers list [flags]
```

### Examples

```
  hostinger vps data-centers list
```

### Options

```
  -h, --help   help for list
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

* [hostinger vps data-centers](hostinger_vps_data-centers.md)	 - Data centers commands

