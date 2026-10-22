## hostinger vps monarx install

Install Monarx

### Synopsis

Install the Monarx malware scanner on a specified virtual machine.

Monarx (https://www.monarx.com/) is a security tool designed to detect and
prevent malware infections on virtual machines. By installing Monarx, users
can enhance the security of their virtual machines, ensuring that they are
protected against malicious software.

Use this endpoint to enable malware protection on VPS instances.

```
hostinger vps monarx install <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps monarx install 1268054
```

### Options

```
  -h, --help                    help for install
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

* [hostinger vps monarx](hostinger_vps_monarx.md)	 - Malware scanner commands

