## hostinger domains whois change-for

Change WHOIS profile for domain

### Synopsis

Change WHOIS contact profile for a domain.

Repoints the given contact roles to a new WHOIS profile and submits the change to the registry.
The profile currently assigned to those roles is resolved automatically;
the request fails if the given roles are not all on the same profile today.

Changing transfer sensitive fields on the owner contact starts an IRTP verification.

The change is processed asynchronously.

Use this endpoint to move a registered domain onto different contact information.

```
hostinger domains whois change-for [flags]
```

### Options

```
      --change-for strings   Contact roles to repoint to the new WHOIS profile (one of: owner, admin, billing, tech)
      --domain string        Domain name
  -h, --help                 help for change-for
      --new-whois-id int     WHOIS profile ID to assign to the domain
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger domains whois](hostinger_domains_whois.md)	 - WHOIS commands

