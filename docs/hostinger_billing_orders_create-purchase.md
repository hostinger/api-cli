## hostinger billing orders create-purchase

Create purchase order

### Synopsis

Create a purchase order for any Hostinger product.

This unified endpoint places an order for one or more catalog items and
works across all Hostinger products, leveraging the existing billing
infrastructure. Use the [catalog endpoint](#tag/billing-catalog) to look
up the `item_id` values available for purchase.

If no payment method is provided, your default payment method will be used automatically.

This endpoint only places the order. Product-specific provisioning
(e.g. VPS setup or domain registration) is not performed here — once the
order completes, use the relevant product endpoints or
[hPanel](https://hpanel.hostinger.com/) to finalize setup.

Use this endpoint to purchase any product available in the catalog.

```
hostinger billing orders create-purchase [flags]
```

### Options

```
      --coupons string          Discount coupon codes (JSON)
  -h, --help                    help for create-purchase
      --items string            Catalog price items to purchase (JSON)
      --payment-method-id int   Payment method ID, default will be used if not provided
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger billing orders](hostinger_billing_orders.md)	 - Orders commands

