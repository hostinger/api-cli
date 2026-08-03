## hostinger vps virtual-machines stop

Stop virtual machine

### Synopsis

Stop a specified virtual machine.

If the virtual machine is already stopped, the request will still be processed without any effect.

This is a compute-only power state change and does not affect billing. To stop future charges,
disable auto-renewal on the owning subscription.

Use this endpoint to power off running VPS instances.

```
hostinger vps virtual-machines stop <virtual-machine-id> [flags]
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

* [hostinger vps virtual-machines](hostinger_vps_virtual-machines.md)	 - Virtual machine commands

