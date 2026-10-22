## hostinger mail webhooks update

Update webhook

### Synopsis

Partially update a webhook. Only the fields included in the request
body are changed; omitted fields retain their current values. Pass
`"description": null` to clear the description.

```
hostinger mail webhooks update <webhook-id> [flags]
```

### Examples

```
  hostinger mail webhooks update 019683f8-1234-7abc-8def-0123456789ab
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
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger mail webhooks](hostinger_mail_webhooks.md)	 - Webhooks commands

