## hostinger mail webhooks regenerate-secret

Regenerate webhook secret

### Synopsis

Regenerate the secret of a webhook. The previous secret is
immediately invalidated. The new secret is returned only in this
response and is sent as a bearer token with every delivery.

```
hostinger mail webhooks regenerate-secret <webhook-id> [flags]
```

### Options

```
  -h, --help   help for regenerate-secret
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger mail webhooks](hostinger_mail_webhooks.md)	 - Webhooks commands

