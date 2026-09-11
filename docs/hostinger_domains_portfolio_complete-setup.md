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

The domain is set up with the default nameservers and without privacy protection. Use
`Update domain nameservers` and `Enable privacy protection` afterwards to change either.

A successful response means the setup request was accepted, not that the domain is already
registered. Poll `Get domain list` for the outcome, the domain appears in `Get domain details`
only once it is registered.

Use this endpoint to finish registering a domain that is awaiting setup on your account.

```
hostinger domains portfolio complete-setup <domain> [flags]
```

### Options

```
      --additional-details string   Additional registration data, possible values depends on TLD (JSON)
      --domain-contacts string      Domain contact information (JSON)
  -h, --help                        help for complete-setup
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger domains portfolio](hostinger_domains_portfolio.md)	 - Portfolio commands

