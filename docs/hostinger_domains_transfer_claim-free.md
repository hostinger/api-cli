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
registrar and can be followed with the [transfer list endpoint](#tag/domains-transfer).

If no WHOIS information is provided, default contact information for that TLD will be used.
Before making request, ensure WHOIS information for desired TLD exists in your account.

Requests which cannot be fulfilled are rejected with an error code in the response body.

Use this endpoint to transfer a domain using a free domain transfer from your account.

```
hostinger domains transfer claim-free [flags]
```

### Options

```
      --auth-code string         Authorization code from the current registrar
      --domain string            Domain name
      --domain-contacts string   Domain contact information (JSON)
  -h, --help                     help for claim-free
      --should-keep-ns           Keep the existing nameservers of the domain (default true)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger domains transfer](hostinger_domains_transfer.md)	 - Transfer commands

