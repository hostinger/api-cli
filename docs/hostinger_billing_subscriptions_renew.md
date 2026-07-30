## hostinger billing subscriptions renew

Renew subscription

### Synopsis

Create a renewal order for an existing Hostinger subscription.

This endpoint places a renewal order for a single subscription, leveraging
the existing billing infrastructure. Use the
[subscriptions endpoint](#tag/billing-subscriptions) to look up the
`subscriptionId` values available for renewal.

If no payment method is provided, your default payment method will be used automatically.

Use this endpoint to renew any subscription available in your account.

```
hostinger billing subscriptions renew <subscription-id> [flags]
```

### Options

```
      --coupons string          Discount coupon codes (JSON)
  -h, --help                    help for renew
      --payment-method-id int   Payment method ID, default will be used if not provided
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger billing subscriptions](hostinger_billing_subscriptions.md)	 - Subscriptions commands

