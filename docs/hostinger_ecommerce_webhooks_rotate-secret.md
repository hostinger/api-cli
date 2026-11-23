## hostinger ecommerce webhooks rotate-secret

Rotate an ecommerce webhook secret

### Synopsis

Explicitly rotate the signing secret of an enabled endpoint on an active store. Save the new secret
securely. Prior secrets retain their original 24-hour overlap deadlines. Repeated calls create
distinct secrets, up to five live secrets. Wait for an old secret to expire at the limit. After a lost
response, explicitly rotate again; stored secrets cannot be retrieved. No daily renewal is required.

409 codes: webhook_endpoint_disabled (endpoint is disabled or revoked), webhook_store_inactive,
webhook_secret_rotation_limit_reached (five live secrets). Errors return the upstream code in
`message`, as `[Ecommerce:<code>] <reason>`.

```
hostinger ecommerce webhooks rotate-secret <store_id> <endpoint_id> [flags]
```

### Examples

```
  hostinger ecommerce webhooks rotate-secret store_01J8Z5F8W9K8M4A7B3C2D1E0FG whep_01J8Z5F8W9K8M4A7B3C2D1E0FG
```

### Options

```
  -h, --help[=true|false]   help for rotate-secret
```

### Options inherited from parent commands

```
      --config string           Config file (default is $HOME/.hostinger.yaml)
      --domain string           Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run[=true|false]    Print the request as a curl command instead of sending it
  -o, --format string           Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string               Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input[=true|false]   Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string         Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string         Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose[=true|false]    Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger ecommerce webhooks](hostinger_ecommerce_webhooks.md)	 - Webhooks commands

