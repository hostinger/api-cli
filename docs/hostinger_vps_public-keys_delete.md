## hostinger vps public-keys delete

Delete public key

### Synopsis

Deprecated: account-level public keys have no direct replacement. Root SSH keys are now managed per
virtual machine via `DELETE /api/vps/v1/virtual-machines/{virtualMachineId}/ssh-keys`.

Delete a public key from your account.

Deleting public key from account does not remove it from virtual machine

Use this endpoint to remove unused SSH keys from account.

```
hostinger vps public-keys delete <public-key-id> [flags]
```

### Examples

```
  hostinger vps public-keys delete 6672861
```

### Options

```
  -h, --help   help for delete
  -y, --yes    Skip the confirmation prompt
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

* [hostinger vps public-keys](hostinger_vps_public-keys.md)	 - Public Keys commands

