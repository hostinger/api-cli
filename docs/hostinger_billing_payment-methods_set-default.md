## hostinger billing payment-methods set-default

Set default payment method

### Synopsis

Set the default payment method for your account.

Use this endpoint to configure the primary payment method for future orders.

```
hostinger billing payment-methods set-default <payment-method-id> [flags]
```

### Examples

```
  hostinger billing payment-methods set-default 9693613
```

### Options

```
  -h, --help   help for set-default
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

* [hostinger billing payment-methods](hostinger_billing_payment-methods.md)	 - Payment methods commands

