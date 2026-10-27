## hostinger vps public-keys attach

Attach public key

### Synopsis

Deprecated: use `POST /api/vps/v1/virtual-machines/{virtualMachineId}/ssh-keys` instead.

Attach existing public keys from your account to a specified virtual machine.

Multiple keys can be attached to a single virtual machine.

Use this endpoint to enable SSH key authentication for VPS instances.

```
hostinger vps public-keys attach <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps public-keys attach 1268054 --ids 18232,10230230
```

### Options

```
  -h, --help                    help for attach
      --ids ints                Public Key IDs to attach
      --wait                    Wait until the action finishes
      --wait-timeout duration   Maximum time to wait with --wait (default 10m0s)
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --domain string     Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger vps public-keys](hostinger_vps_public-keys.md)	 - Public Keys commands

