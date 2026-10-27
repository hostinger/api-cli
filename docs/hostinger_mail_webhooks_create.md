## hostinger mail webhooks create

Create webhook

### Synopsis

Create a webhook for the given mailbox. The generated secret is
returned only in this response and is sent as a bearer token with
every delivery.

```
hostinger mail webhooks create <mailbox-id> [flags]
```

### Examples

```
  hostinger mail webhooks create AC1a2b3c4d5e6f7g --events message.received --name 'New message notifier' --url https://example.com/webhooks/incoming
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
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --domain string     Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger mail webhooks](hostinger_mail_webhooks.md)	 - Webhooks commands

