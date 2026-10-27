## hostinger ecommerce discounts create

Create a discount

### Synopsis

Create a discount for a store. Fixed discounts take an amount in the smallest currency
unit (e.g. $10 is 1000); percentage discounts take a whole-number value between 1 and 100.
Free-shipping discounts ignore value. Returns the created discount.

```
hostinger ecommerce discounts create <store_id> [flags]
```

### Examples

```
  hostinger ecommerce discounts create store_01J8Z5F8W9K8M4A7B3C2D1E0FG --code BLACKFRIDAY --type percentage --value 20
```

### Options

```
      --allocation string    Whether the discount applies to the cart total or to each eligible item. (one of: total, item)
      --code string          The discount code customers enter at checkout.
      --ends-at string       When the discount expires. A bare date runs to the end of that day in time_zone. Never expires when omitted.
  -h, --help                 help for create
      --min-cart-value int   Minimum cart value in the smallest currency unit required for the discount to apply.
      --name string          A human-friendly discount name.
      --starts-at string     When the discount becomes active. A bare date (2026-11-27) anchors to time_zone. Defaults to now when omitted.
      --time-zone string     IANA time zone used to interpret starts_at and ends_at.
      --type string          The discount type. (one of: percentage, fixed, free_shipping)
      --usage-limit int      Maximum number of times the discount can be redeemed.
      --value int            For percentage discounts a whole number 1-100; for fixed discounts an amount in the smallest currency unit (e.g. $10 is 1000). Ignored for free_shipping.
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

* [hostinger ecommerce discounts](hostinger_ecommerce_discounts.md)	 - Discounts commands

