## hostinger domains portfolio purchase

Purchase new domain

### Synopsis

Purchase and register a new domain name.

If registration fails, login to hPanel (https://hpanel.hostinger.com/) and check domain registration status.

If no payment method is provided, your default payment method will be used automatically.

If the response is `202 Accepted`, the payment is still being processed and the domain was
not registered. Once the order completes, register the domain from
hPanel (https://hpanel.hostinger.com/).

If no WHOIS information is provided, default contact information for that TLD will be used.
Before making request, ensure WHOIS information for desired TLD exists in your account.

Some TLDs require `additional_details` to be provided and these will be validated before completing purchase.

Use this endpoint to register new domains for users.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --additional-details
    {}

  --coupons
    [
      "Coupon 3"
    ]

  --domain-contacts
    {
      "admin_id": 546123,
      "billing_id": 741288,
      "owner_id": 741288,
      "tech_id": 741288
    }

```
hostinger domains portfolio purchase [flags]
```

### Examples

```
  hostinger domains portfolio purchase --domain my-new-domain.tld --item-id hostingercom-domain-com-usd-1y
```

### Options

```
      --additional-details string   Additional registration data, possible values depends on TLD (JSON, @file or -; shape in the help above)
      --coupons string              Discount coupon codes (JSON, @file or -; shape in the help above)
      --domain string               Domain name
      --domain-contacts string      Domain contact information (JSON, @file or -; shape in the help above)
  -h, --help                        help for purchase
      --item-id string              Catalog price item ID
      --payment-method-id int       Payment method ID, default will be used if not provided
  -y, --yes                         Skip the confirmation prompt
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

* [hostinger domains portfolio](hostinger_domains_portfolio.md)	 - Portfolio commands

