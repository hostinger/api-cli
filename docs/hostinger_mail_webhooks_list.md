## hostinger mail webhooks list

List webhooks

### Synopsis

Retrieve a paginated list of webhooks belonging to the given mail
order. Supports filtering by mailbox and status. The webhook secret
is never included; it is returned only when a webhook is created or
its secret is regenerated.

```
hostinger mail webhooks list <order-id> [flags]
```

### Options

```
  -h, --help                help for list
      --mailbox-id string   Filter by the mailbox resource ID the webhooks are attached to
      --page int            Page number
      --per-page int        Number of items per page (default 25)
      --status string       Filter webhooks by status (one of: active, disabled, paused)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger mail webhooks](hostinger_mail_webhooks.md)	 - Webhooks commands

