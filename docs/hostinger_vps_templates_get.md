## hostinger vps templates get

Get template details

### Synopsis

Retrieve detailed information about a specific OS template for virtual machines.

Use this endpoint to view specific template specifications before deployment.

```
hostinger vps templates get <template-id> [flags]
```

### Examples

```
  hostinger vps templates get 2868928
```

### Options

```
  -h, --help   help for get
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

* [hostinger vps templates](hostinger_vps_templates.md)	 - OS Templates commands

