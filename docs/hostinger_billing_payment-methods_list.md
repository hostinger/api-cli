## hostinger billing payment-methods list

Get payment method list

### Synopsis

Retrieve available payment methods that can be used for placing new orders.

If you want to add new payment method,
please use hPanel (https://hpanel.hostinger.com/billing/payment-methods).

Use this endpoint to view available payment options before creating orders.

```
hostinger billing payment-methods list [flags]
```

### Examples

```
  hostinger billing payment-methods list
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

* [hostinger billing payment-methods](hostinger_billing_payment-methods.md)	 - Payment methods commands

