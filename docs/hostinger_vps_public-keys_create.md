## hostinger vps public-keys create

Create public key

### Synopsis

Deprecated: account-level public keys have no direct replacement. Root SSH keys are now managed per
virtual machine via `POST /api/vps/v1/virtual-machines/{virtualMachineId}/ssh-keys`.

Add a new public key to your account.

Use this endpoint to register SSH keys for VPS authentication.

```
hostinger vps public-keys create [flags]
```

### Options

```
  -h, --help          help for create
      --key string    
      --name string   
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger vps public-keys](hostinger_vps_public-keys.md)	 - Public Keys commands

