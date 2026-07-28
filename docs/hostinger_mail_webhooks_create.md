## hostinger mail webhooks create

Create webhook

### Synopsis

Create a webhook for the given mailbox. The generated secret is
returned only in this response and is sent as a bearer token with
every delivery.

```
hostinger mail webhooks create <mailbox-id> [flags]
```

### Options

```
      --description string   Optional description of the webhook's purpose
      --events strings       Events that trigger this webhook (one of: message.received)
  -h, --help                 help for create
      --name string          Human-readable name for this webhook
      --status string        Initial status of the webhook (one of: active, disabled, paused) (default "active")
      --url string           Publicly reachable URL that receives the webhook POST requests
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger mail webhooks](hostinger_mail_webhooks.md)	 - Webhooks commands

