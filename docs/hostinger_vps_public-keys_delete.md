## hostinger vps public-keys delete

Delete public key

### Synopsis

Deprecated: account-level public keys have no direct replacement. Root SSH keys are now managed per
virtual machine via `DELETE /api/vps/v1/virtual-machines/{virtualMachineId}/ssh-keys`.

Delete a public key from your account. 

**Deleting public key from account does not remove it from virtual machine** 
       
Use this endpoint to remove unused SSH keys from account.

```
hostinger vps public-keys delete <public-key-id> [flags]
```

### Options

```
  -h, --help   help for delete
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger vps public-keys](hostinger_vps_public-keys.md)	 - Public Keys commands

