## hostinger mail forwarders create

Create forwarder

### Synopsis

Create a forwarder from the given mailbox to the destination address.
The destination receives a confirmation email and forwarding becomes
active only after it is confirmed.

```
hostinger mail forwarders create <mailbox-id> [flags]
```

### Options

```
      --destination string     Email address the messages will be forwarded to
  -h, --help                   help for create
      --is-keep-copy-enabled   Whether to keep a copy of forwarded messages in the mailbox. Defaults to false.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger mail forwarders](hostinger_mail_forwarders.md)	 - Forwarders commands

