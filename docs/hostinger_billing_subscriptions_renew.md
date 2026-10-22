## hostinger billing subscriptions renew

Renew subscription

### Synopsis

Create a renewal order for an existing Hostinger subscription.

This endpoint places a renewal order for a single subscription, leveraging
the existing billing infrastructure. Use the
subscriptions endpoint (#tag/billing-subscriptions) to look up the
`subscriptionId` values available for renewal.

If no payment method is provided, your default payment method will be used automatically.

If the response is `202 Accepted`, the payment is still being processed and the renewal will
complete asynchronously once the payment is confirmed.

Use this endpoint to renew any subscription available in your account.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --coupons
    [
      "Coupon 3"
    ]

```
hostinger billing subscriptions renew <subscription-id> [flags]
```

### Examples

```
  hostinger billing subscriptions renew Cxy353Uhl1xC54pG6
```

### Options

```
      --coupons string          Discount coupon codes (JSON, @file or -; shape in the help above)
  -h, --help                    help for renew
      --payment-method-id int   Payment method ID, default will be used if not provided
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

