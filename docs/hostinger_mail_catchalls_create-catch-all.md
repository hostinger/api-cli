## hostinger mail catchalls create-catch-all

Create catch-all

### Synopsis

Create a catch-all that routes all messages sent to unknown addresses
of the domain to the given mailbox. The mailbox address receives a
confirmation email and the catch-all becomes active only after it is
confirmed. A domain can have only one catch-all.

```
hostinger mail catchalls create-catch-all <mailbox-id> [flags]
```

### Options

```
  -h, --help   help for create-catch-all
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger mail catchalls](hostinger_mail_catchalls.md)	 - Catchalls commands

