## hostinger domains move reject-incoming

Reject incoming domain move

### Synopsis

Reject an incoming move for a specified domain.

The domain stays in the account which initiated the move.
Moves you have already accepted cannot be rejected anymore.

Use this endpoint to decline a domain you do not want to take over.

```
hostinger domains move reject-incoming <domain> [flags]
```

### Options

```
  -h, --help   help for reject-incoming
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger domains move](hostinger_domains_move.md)	 - Move commands

