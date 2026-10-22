## hostinger billing subscriptions disable-auto-renewal

Disable auto-renewal

### Synopsis

Disable auto-renewal for a subscription.

Use this endpoint when disable auto-renewal for a subscription.

```
hostinger billing subscriptions disable-auto-renewal <subscription-id> [flags]
```

### Examples

```
  hostinger billing subscriptions disable-auto-renewal Cxy353Uhl1xC54pG6
```

### Options

```
  -h, --help   help for disable-auto-renewal
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

* [hostinger billing subscriptions](hostinger_billing_subscriptions.md)	 - Subscriptions commands

