## hapi vps public-keys attach

Attach public key

### Synopsis

Attach existing public keys from your account to a specified virtual machine.

Multiple keys can be attached to a single virtual machine.

Use this endpoint to enable SSH key authentication for VPS instances.

```
hapi vps public-keys attach <virtual-machine-id> [flags]
```

### Options

```
  -h, --help       help for attach
      --ids ints   Public Key IDs to attach
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps public-keys](hapi_vps_public-keys.md)	 - Public Keys commands

