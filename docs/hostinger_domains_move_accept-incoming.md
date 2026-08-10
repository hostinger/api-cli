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
[incoming move endpoint](#tag/domains-move).

Use this endpoint to take ownership of a domain offered to you.

```
hostinger domains move accept-incoming <domain> [flags]
```

### Options

```
      --domain-contacts string   WHOIS profiles of the accepting account. Only the contact types required by the TLD are applied, but all four IDs must be provided. (JSON)
  -h, --help                     help for accept-incoming
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger domains move](hostinger_domains_move.md)	 - Move commands

