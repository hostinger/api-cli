## hostinger domains move start-outgoing

Start outgoing domain move

### Synopsis

Initiate a move of a specified domain to another Hostinger account.

The receiving account has to already exist and accept the move before the domain changes hands.

The domain must be active. The subscription it belongs to is resolved automatically,
and the request is rejected with a 404 status code when the domain has no domain
subscription of its own.

Domains protected by premium protection require an additional verification step,
such requests are rejected with a 428 status code.

Use this endpoint to hand a domain over to another Hostinger user.

```
hostinger domains move start-outgoing <domain> [flags]
```

### Options

```
  -h, --help                        help for start-outgoing
      --new-customer-email string   Email address of the Hostinger account receiving the domain
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger domains move](hostinger_domains_move.md)	 - Move commands

