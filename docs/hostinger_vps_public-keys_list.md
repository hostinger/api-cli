## hostinger vps public-keys list

Get public keys

### Synopsis

Deprecated: account-level public keys have no direct replacement. Root SSH keys are now managed per
virtual machine via `GET /api/vps/v1/virtual-machines/{virtualMachineId}/ssh-keys`.

Retrieve public keys associated with your account.

Use this endpoint to view available SSH keys for VPS authentication.

```
hostinger vps public-keys list [flags]
```

### Options

```
  -h, --help       help for list
      --page int   Page number
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger vps public-keys](hostinger_vps_public-keys.md)	 - Public Keys commands

