## hostinger domains portfolio claim-free

Claim free domain

### Synopsis

Claim a free domain available on your account and register it.

Unlike purchasing a domain, this consumes a free domain you already have,
so no payment method is required.

A successful response means the domain is registered. If registration fails, login to
[hPanel](https://hpanel.hostinger.com/) and check domain registration status.

If no WHOIS information is provided, default contact information for that TLD will be used.
Before making request, ensure WHOIS information for desired TLD exists in your account.

Some TLDs require `additional_details` to be provided and these will be validated before claiming.

Requests which cannot be fulfilled are rejected with an error code in the response body,
for example `2037` when no free domain is available.

Use this endpoint to register a domain using a free domain from your account.

```
hostinger domains portfolio claim-free [flags]
```

### Options

```
      --additional-details string   Additional registration data, possible values depends on TLD (JSON)
      --domain string               Domain name
      --domain-contacts string      Domain contact information (JSON)
  -h, --help                        help for claim-free
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger domains portfolio](hostinger_domains_portfolio.md)	 - Portfolio commands

