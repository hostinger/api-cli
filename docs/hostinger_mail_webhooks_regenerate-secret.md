## hostinger mail webhooks regenerate-secret

Regenerate webhook secret

### Synopsis

Regenerate the secret of a webhook. The previous secret is
immediately invalidated. The new secret is returned only in this
response and is sent as a bearer token with every delivery.

```
hostinger mail webhooks regenerate-secret <webhook-id> [flags]
```

### Examples

```
  hostinger mail webhooks regenerate-secret 019683f8-1234-7abc-8def-0123456789ab
```

### Options

```
  -h, --help   help for regenerate-secret
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

