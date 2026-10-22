## hostinger vps ptr create

Create PTR record

### Synopsis

Create or update a PTR (Pointer) record for a specified virtual machine.

Use this endpoint to configure reverse DNS lookup for VPS IP addresses.

```
hostinger vps ptr create <virtual-machine-id> <ip-address-id> [flags]
```

### Examples

```
  hostinger vps ptr create 1268054 246547 --domain example.tld
```

### Options

```
      --domain string           Pointer record domain
  -h, --help                    help for create
      --wait                    Wait until the action finishes
      --wait-timeout duration   Maximum time to wait with --wait (default 10m0s)
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

* [hostinger vps ptr](hostinger_vps_ptr.md)	 - PTR records commands

