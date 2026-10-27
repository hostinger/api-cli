## hostinger mail orders list

List orders

### Synopsis

Retrieve a paginated list of mail orders associated with your account.

Use this endpoint to monitor your mail services, including their status,
plan, attached domain, and expiration details.

```
hostinger mail orders list [flags]
```

### Examples

```
  hostinger mail orders list
```

### Options

```
      --all             Fetch every page and print all items
      --domain string   Filter orders by domain name (exact match)
  -h, --help            help for list
      --is-trial        Filter orders by trial state
      --page int        Page number
      --per-page int    Number of items per page (default 25)
      --sort -          Sort orders by field. Prefix with - for descending order. (one of: created_at, -created_at, expires_at, -expires_at) (default "-created_at")
      --status string   Filter orders by status (one of: pending_setup, active, suspended)
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger mail orders](hostinger_mail_orders.md)	 - Orders commands

