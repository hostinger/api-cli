## hostinger mail api-tokens create

Create API token

### Synopsis

Create an API token for the given mail order. The token grants access
to the [Hostinger Email API](https://api.mail.hostinger.com/), where
you can provision and manage the mailboxes it is scoped to.

The plaintext token is returned only in this response, never again.
A maximum of 10 tokens can exist per order. Use
`scope.has_all_mailboxes` to cover all current and future mailboxes,
or list specific mailboxes in `scope.mailbox_ids`.

```
hostinger mail api-tokens create <order-id> [flags]
```

### Options

```
  -h, --help           help for create
      --name string    Human-readable label for this token
      --scope string   Mailbox scope this token can access (JSON)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger mail api-tokens](hostinger_mail_api-tokens.md)	 - API Tokens commands

