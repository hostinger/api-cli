## hostinger mail webhooks update

Update webhook

### Synopsis

Partially update a webhook. Only the fields included in the request
body are changed; omitted fields retain their current values. Pass
`"description": null` to clear the description.

```
hostinger mail webhooks update <webhook-id> [flags]
```

### Options

```
      --description string   New description, or null to clear it
      --events strings       Replaces the full list of subscribed events (one of: message.received)
  -h, --help                 help for update
      --name string          New human-readable name for the webhook
      --status string        New status for the webhook (one of: active, disabled, paused)
      --url string           New URL to deliver events to
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger mail webhooks](hostinger_mail_webhooks.md)	 - Webhooks commands

