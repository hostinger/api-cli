## hostinger domains transfer start

Start domain transfer

### Synopsis

Transfer a domain from another registrar to your account.

The transfer runs on a domain transfer service you have already purchased.

Before making request, unlock the domain at the current registrar and, if its TLD uses one,
get its authorization code.

A successful response means the transfer has been started. Completion depends on the current
registrar and can be followed with the transfer list endpoint (#tag/domains-transfer).

If no WHOIS information is provided, default contact information for that TLD will be used.
Before making request, ensure WHOIS information for desired TLD exists in your account.

Use this endpoint to bring domains registered elsewhere into your account.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --domain-contacts
    {
      "admin_id": 546123,
      "billing_id": 741288,
      "owner_id": 741288,
      "tech_id": 741288
    }

```
hostinger domains transfer start [flags]
```

### Examples

```
  hostinger domains transfer start --domain my-new-domain.tld
```

### Options

```
      --auth-code string              EPP code that authorizes this transfer, which the domain owner gets from the current registrar. Required only for TLDs whose transfers use one.
      --domain string                 Domain name
      --domain-contacts string        Domain contact information (JSON, @file or -; shape in the help above)
  -h, --help[=true|false]             help for start
      --should-keep-ns[=true|false]   Keep the existing nameservers of the domain (default true)
```

### Options inherited from parent commands

```
      --config string           Config file (default is $HOME/.hostinger.yaml)
      --dry-run[=true|false]    Print the request as a curl command instead of sending it
  -o, --format string           Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string               Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input[=true|false]   Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string         Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string         Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose[=true|false]    Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger domains transfer](hostinger_domains_transfer.md)	 - Transfer commands

