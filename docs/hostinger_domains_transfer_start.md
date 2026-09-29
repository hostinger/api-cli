## hostinger domains transfer start

Start domain transfer

### Synopsis

Transfer a domain from another registrar to your account.

The transfer runs on a domain transfer service you have already purchased.

Before making request, unlock the domain at the current registrar and get its authorization
code.

A successful response means the transfer has been started. Completion depends on the current
registrar and can be followed with the [transfer list endpoint](#tag/domains-transfer).

If no WHOIS information is provided, default contact information for that TLD will be used.
Before making request, ensure WHOIS information for desired TLD exists in your account.

Use this endpoint to bring domains registered elsewhere into your account.

```
hostinger domains transfer start [flags]
```

### Options

```
      --auth-code string         Authorization code from the current registrar
      --domain string            Domain name
      --domain-contacts string   Domain contact information (JSON)
  -h, --help                     help for start
      --should-keep-ns           Keep the existing nameservers of the domain (default true)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger domains transfer](hostinger_domains_transfer.md)	 - Transfer commands

