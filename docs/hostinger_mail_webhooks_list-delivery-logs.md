## hostinger mail webhooks list-delivery-logs

List webhook delivery logs

### Synopsis

Retrieve a paginated list of webhook delivery logs for the given mail
order, including delivery outcome, duration, and retry counts.
Supports filtering by mailbox.

```
hostinger mail webhooks list-delivery-logs <order-id> [flags]
```

### Options

```
  -h, --help                help for list-delivery-logs
      --mailbox-id string   Filter by the mailbox resource ID the webhooks are attached to
      --page int            Page number
      --per-page int        Number of items per page (default 25)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger mail webhooks](hostinger_mail_webhooks.md)	 - Webhooks commands

