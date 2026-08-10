## hostinger domains move incoming

Get incoming domain move

### Synopsis

Retrieve the incoming move for a specified domain.

Returns 404 when no account is moving this domain to you.

Use this endpoint to check whether a domain addressed to you is still waiting to be accepted.

```
hostinger domains move incoming <domain> [flags]
```

### Options

```
      --force-sync activating   Re-check the move against the registry before responding. Only has an effect while the move is in the activating status.
  -h, --help                    help for incoming
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger domains move](hostinger_domains_move.md)	 - Move commands

