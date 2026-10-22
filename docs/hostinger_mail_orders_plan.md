## hostinger mail orders plan

Get order plan

### Synopsis

Retrieve the plan the given mail order was purchased with, including
domain-level and mailbox-level quotas, limits, and protocol
availability.

```
hostinger mail orders plan <order-id> [flags]
```

### Examples

```
  hostinger mail orders plan OR1a2b3c4d5e6f7g
```

### Options

```
  -h, --help   help for plan
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

* [hostinger mail orders](hostinger_mail_orders.md)	 - Orders commands

