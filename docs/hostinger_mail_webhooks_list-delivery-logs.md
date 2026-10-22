## hostinger mail webhooks list-delivery-logs

List webhook delivery logs

### Synopsis

Retrieve a paginated list of webhook delivery logs for the given mail
order, including delivery outcome, duration, and retry counts.
Supports filtering by mailbox.

```
hostinger mail webhooks list-delivery-logs <order-id> [flags]
```

### Examples

```
  hostinger mail webhooks list-delivery-logs OR1a2b3c4d5e6f7g
```

### Options

```
      --all                 Fetch every page and print all items
  -h, --help                help for list-delivery-logs
      --mailbox-id string   Filter by the mailbox resource ID the webhooks are attached to
      --page int            Page number
      --per-page int        Number of items per page (default 25)
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

