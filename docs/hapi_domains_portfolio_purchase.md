## hapi domains portfolio purchase

Purchase new domain

### Synopsis

Purchase and register a new domain name.

If registration fails, login to [hPanel](https://hpanel.hostinger.com/) and check domain registration status.

If no payment method is provided, your default payment method will be used automatically.

If no WHOIS information is provided, default contact information for that TLD will be used.
Before making request, ensure WHOIS information for desired TLD exists in your account.

Some TLDs require `additional_details` to be provided and these will be validated before completing purchase.

Use this endpoint to register new domains for users.

```
hapi domains portfolio purchase [flags]
```

### Options

```
      --additional-details string   Additional registration data, possible values depends on TLD (JSON)
      --coupons string              Discount coupon codes (JSON)
      --domain string               Domain name
      --domain-contacts string      Domain contact information (JSON)
  -h, --help                        help for purchase
      --item-id string              Catalog price item ID
      --payment-method-id int       Payment method ID, default will be used if not provided
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi domains portfolio](hapi_domains_portfolio.md)	 - Portfolio commands

