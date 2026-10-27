## hostinger domains transfer claim-free

Claim free domain transfer

### Synopsis

Claim a free domain transfer available on your account and start the transfer.

Unlike purchasing a transfer, this consumes a free domain transfer you already have,
so no payment method is required.

Before making request, unlock the domain at the current registrar and get its authorization
code. The transfer is validated first, so domains which cannot be transferred are rejected
before the free domain transfer is consumed.

A successful response means the transfer has been started. Completion depends on the current
registrar and can be followed with the transfer list endpoint (#tag/domains-transfer).

If no WHOIS information is provided, default contact information for that TLD will be used.
Before making request, ensure WHOIS information for desired TLD exists in your account.

Requests which cannot be fulfilled are rejected with an error code in the response body.

Use this endpoint to transfer a domain using a free domain transfer from your account.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --domain-contacts
    {
      "admin_id": 546123,
      "billing_id": 741288,
      "owner_id": 741288,
      "tech_id": 741288
    }

```
hostinger domains transfer claim-free [flags]
```

### Examples

```
  hostinger domains transfer claim-free --auth-code Auth123Code456 --domain my-new-domain.tld
```

### Options

```
      --auth-code string         Authorization code from the current registrar
      --domain string            Domain name
      --domain-contacts string   Domain contact information (JSON, @file or -; shape in the help above)
  -h, --help                     help for claim-free
      --should-keep-ns           Keep the existing nameservers of the domain (default true)
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger domains transfer](hostinger_domains_transfer.md)	 - Transfer commands

