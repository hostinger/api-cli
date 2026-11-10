## hostinger domains whois create

Create WHOIS profile

### Synopsis

Create WHOIS contact profile.

Use this endpoint to add new contact information for domain registration.

The keys required in `whois_details` and `tld_details` depend on the TLD, country and entity type, and can be
country-specific: for Lithuania, for example, `state_lt` and `zip_lt` are required instead of `state` and
`zip`. A 422 response lists every required key that is missing or has an invalid value under
`errors.whois_details`, and under `errors.tld_details` when `tld_details` is sent. Each error states the
key's constraint and, where the key takes a fixed set of values, the allowed values, so all of them can be
corrected in one retry.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --tld-details
    {}

  --whois-details
    {}

```
hostinger domains whois create [flags]
```

### Examples

```
  hostinger domains whois create --country NL --entity-type individual --tld com --whois-details '{}'
```

### Options

```
      --country string           ISO 3166 2-letter country code
      --entity-type string       Legal entity type (one of: individual, organization)
  -h, --help                     help for create
      --tld string               TLD of the domain (without leading dot)
      --tld-details string       TLD-specific details. The required keys depend on the TLD, country and entity type. (JSON, @file or -; shape in the help above)
      --whois-details state_lt   WHOIS contact details. The required keys depend on the TLD, country and entity type, and can be country-specific (e.g. state_lt and `zip_lt` instead of `state` and `zip` for Lithuania). (JSON, @file or -; shape in the help above)
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

* [hostinger domains whois](hostinger_domains_whois.md)	 - WHOIS commands

