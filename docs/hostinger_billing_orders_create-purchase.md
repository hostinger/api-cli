## hostinger billing orders create-purchase

Create purchase order

### Synopsis

Create a purchase order for any Hostinger product.

This unified endpoint places an order for one or more catalog items and
works across all Hostinger products, leveraging the existing billing
infrastructure. Use the catalog endpoint (#tag/billing-catalog) to look
up the `item_id` values available for purchase.

If no payment method is provided, your default payment method will be used automatically.

If the response is `202 Accepted`, the payment is still being processed and the order will
complete asynchronously once the payment is confirmed.

This endpoint only places the order. Product-specific provisioning
(e.g. VPS setup or domain registration) is not performed here — once the
order completes, use the relevant product endpoints or
hPanel (https://hpanel.hostinger.com/) to finalize setup.

Use this endpoint to purchase any product available in the catalog.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --coupons
    [
      "Coupon 3"
    ]

  --items
    [
      {
        "item_id": "hostingercom-vps-kvm2-usd-1m",
        "quantity": 1
      }
    ]

```
hostinger billing orders create-purchase [flags]
```

### Examples

```
  hostinger billing orders create-purchase --items '[{"item_id":"hostingercom-vps-kvm2-usd-1m","quantity":1}]'
```

### Options

```
      --coupons string          Discount coupon codes (JSON, @file or -; shape in the help above)
  -h, --help                    help for create-purchase
      --items string            Catalog price items to purchase (JSON, @file or -; shape in the help above)
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

* [hostinger billing orders](hostinger_billing_orders.md)	 - Orders commands

