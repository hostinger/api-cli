## hostinger domains verifications direct

Get domain verifications

### Synopsis

Retrieve a list of pending and completed domain verifications.

```
hostinger domains verifications direct [flags]
```

### Examples

```
  hostinger domains verifications direct --domains value
```

### Options

```
      --domains strings   The list of domains for which to get verification details for.
  -h, --help              help for direct
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

* [hostinger domains verifications](hostinger_domains_verifications.md)	 - Verifications commands

