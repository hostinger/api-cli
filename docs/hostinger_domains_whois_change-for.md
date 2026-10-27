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

### Examples

```
  hostinger domains whois change-for --change-for owner,admin --domain mydomain.tld --new-whois-id 131502
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

* [hostinger domains whois](hostinger_domains_whois.md)	 - WHOIS commands

