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

### Options

```
  -h, --help                   help for start
      --root-password string   Temporary root password for recovery mode
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger vps recovery](hostinger_vps_recovery.md)	 - Recovery commands

