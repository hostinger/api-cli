## hostinger mail orders list

Get mail order list

### Synopsis

Retrieve a paginated list of mail orders associated with your account.

Use this endpoint to monitor your mail services, including their status,
plan, attached domain, and expiration details.

```
hostinger mail orders list [flags]
```

### Options

```
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
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger mail orders](hostinger_mail_orders.md)	 - Orders commands

