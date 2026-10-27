## hostinger domains portfolio complete-setup

Complete domain setup

### Synopsis

Register a domain you have already paid for but which has not been set up yet.

Use this endpoint when an order completed without registering the domain, for example when
`Purchase new domain` returned `202 Accepted` and the domain was added to your account without
being registered, or when an earlier setup attempt failed. No new order is placed and no payment
is taken: the subscription you already own is used, for the period you already paid for.

A domain is left awaiting setup when the details needed to register it were missing or invalid
as the order completed. Domains ordered elsewhere can be awaiting setup for the same reason.
Complete the missing information, then call this endpoint. If the order itself has not completed
yet, the domain is not on your account, wait until it appears in `Get domain list`.

If `domain_contacts` is omitted, the default WHOIS profile of that TLD is used for all four
roles. The profile must exist and be complete for the TLD, an incomplete profile is the most
common reason a domain is left awaiting setup. Create one with `Create WHOIS profile`.

Some TLDs require `additional_details`. These are validated before setup, so a missing or
invalid value is rejected without any registration being attempted.

The domain is set up with the default nameservers and with privacy protection. Use
`Update domain nameservers` and `Disable privacy protection` afterwards to change either.

A successful response means the setup request was accepted, not that the domain is already
registered. Poll `Get domain list` for the outcome, the domain appears in `Get domain details`
only once it is registered.

Use this endpoint to finish registering a domain that is awaiting setup on your account.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --additional-details
    {}

  --domain-contacts
    {
      "admin_id": 546123,
      "billing_id": 741288,
      "owner_id": 741288,
      "tech_id": 741288
    }

```
hostinger domains portfolio complete-setup <domain> [flags]
```

### Examples

```
  hostinger domains portfolio complete-setup mydomain.tld
```

### Options

```
      --additional-details string   Additional registration data, possible values depends on TLD (JSON, @file or -; shape in the help above)
      --domain-contacts string      Domain contact information (JSON, @file or -; shape in the help above)
  -h, --help                        help for complete-setup
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

