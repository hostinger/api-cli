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

### Examples

```
  hostinger mail webhooks list OR1a2b3c4d5e6f7g
```

### Options

```
      --all                 Fetch every page and print all items
  -h, --help                help for list
      --mailbox-id string   Filter by the mailbox resource ID the webhooks are attached to
      --page int            Page number
      --per-page int        Number of items per page (default 25)
      --status string       Filter webhooks by status (one of: active, disabled, paused)
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

