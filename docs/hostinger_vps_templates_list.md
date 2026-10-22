## hostinger vps templates list

Get templates

### Synopsis

Retrieve available OS templates for virtual machines.

Use this endpoint to view operating system options before creating or recreating VPS instances.

```
hostinger vps templates list [flags]
```

### Examples

```
  hostinger vps templates list
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

* [hostinger vps templates](hostinger_vps_templates.md)	 - OS Templates commands

