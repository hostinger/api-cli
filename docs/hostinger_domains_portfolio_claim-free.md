## hostinger domains portfolio claim-free

Claim free domain

### Synopsis

Claim a free domain available on your account and register it.

Unlike purchasing a domain, this consumes a free domain you already have,
so no payment method is required.

A successful response means the domain is registered. If registration fails, login to
hPanel (https://hpanel.hostinger.com/) and check domain registration status.

If no WHOIS information is provided, default contact information for that TLD will be used.
Before making request, ensure WHOIS information for desired TLD exists in your account.

Some TLDs require `additional_details` to be provided and these will be validated before claiming.

Requests which cannot be fulfilled are rejected with an error code in the response body,
for example `2037` when no free domain is available.

Use this endpoint to register a domain using a free domain from your account.

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
hostinger domains portfolio claim-free [flags]
```

### Examples

```
  hostinger domains portfolio claim-free --domain my-new-domain.tld
```

### Options

```
      --additional-details string   Additional registration data, possible values depends on TLD (JSON, @file or -; shape in the help above)
      --domain string               Domain name
      --domain-contacts string      Domain contact information (JSON, @file or -; shape in the help above)
  -h, --help                        help for claim-free
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

