## hostinger vps recovery start

Start recovery mode

### Synopsis

Initiate recovery mode for a specified virtual machine.

Recovery mode is a special state that allows users to perform system rescue operations,
such as repairing file systems, recovering data, or troubleshooting issues that prevent the virtual machine
from booting normally.

Virtual machine will boot recovery disk image and original disk image will be mounted in `/mnt` directory.

Use this endpoint to enable system rescue operations on VPS instances.

```
hostinger vps recovery start <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps recovery start 1268054 --root-password oMeNRustosIO
```

### Options

```
  -h, --help                    help for start
      --root-password string    Temporary root password for recovery mode
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

* [hostinger vps recovery](hostinger_vps_recovery.md)	 - Recovery commands

