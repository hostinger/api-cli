## hostinger billing subscriptions list

Get subscription list

### Synopsis

Retrieve a list of all subscriptions associated with your account.

Use this endpoint to monitor active services and billing status.

```
hostinger billing subscriptions list [flags]
```

### Examples

```
  hostinger billing subscriptions list
```

### Options

```
  -h, --help   help for list
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

