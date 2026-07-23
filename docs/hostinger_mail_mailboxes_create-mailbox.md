## hostinger mail mailboxes create-mailbox

Create mailbox

### Synopsis

Create a mailbox under the given mail order. The full email address is
composed from the given local part and the domain of the order.

```
hostinger mail mailboxes create-mailbox <order-id> [flags]
```

### Options

```
  -h, --help                help for create-mailbox
      --local-part string   Local part of the mailbox address (the part before the @). The domain is taken from the order. Must start and end with a letter or digit; single dots, underscores and hyphens are allowed in between.
      --password string     Mailbox password. Minimum 8 characters with uppercase, lowercase, number and special character.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger mail mailboxes](hostinger_mail_mailboxes.md)	 - Mailboxes commands

