## hostinger domains move accept-incoming

Accept incoming domain move

### Synopsis

Accept an incoming move for a specified domain.

The provided WHOIS profiles become the contacts of the domain, so they must belong
to your account and satisfy the requirements of the TLD. Only the contact types the
domain actually uses are applied, but all four profile IDs have to be provided.

The move has to still be waiting for your decision, already accepted moves
cannot be accepted again.

Accepting does not complete the move. A confirmation email is sent to the email address of
the new owner contact, and the domain changes hands only after the change is confirmed from it.
Until then the move stays in the `activating` status, which can be followed with the
incoming move endpoint (#tag/domains-move).

Use this endpoint to take ownership of a domain offered to you.

The domain can be given as [domain] or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link` when that website's domain belongs to the account, default_domain in the config file, or a prompt when run in a terminal.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --domain-contacts
    {
      "admin_id": 114698,
      "billing_id": 154698,
      "owner_id": 614698,
      "tech_id": 524248
    }

```
hostinger domains move accept-incoming [domain] [flags]
```

### Examples

```
  hostinger domains move accept-incoming mydomain.tld --domain-contacts @domain-contacts.json
```

### Options

```
      --domain-contacts string   WHOIS profiles of the accepting account. Only the contact types required by the TLD are applied, but all four IDs must be provided. (JSON, @file or -; shape in the help above)
  -h, --help                     help for accept-incoming
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

* [hostinger domains move](hostinger_domains_move.md)	 - Move commands

