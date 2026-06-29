## hostinger vps recovery stop

Stop recovery mode

### Synopsis

Stop recovery mode for a specified virtual machine.

If virtual machine is not in recovery mode, this operation will fail.

Use this endpoint to exit system rescue mode and return VPS to normal operation.

```
hostinger vps recovery stop <virtual-machine-id> [flags]
```

### Options

```
  -h, --help   help for stop
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger vps recovery](hostinger_vps_recovery.md)	 - Recovery commands

