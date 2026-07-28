## hostinger mail logs list-access

List access logs

### Synopsis

Retrieve paginated access logs for the domain attached to the given
mail order. Supports filtering by account, date range, protocol,
status, and deletion flag. Results are sorted by timestamp descending.

```
hostinger mail logs list-access <order-id> [flags]
```

### Options

```
      --account string     Filter log entries by a specific email account
      --date from_date     Exact date filter (YYYY-MM-DD). Takes precedence over from_date/`to_date` when both are given.
      --from-date string   Date range start (RFC 3339)
      --has-deletions      Filter access log entries by whether the session had deletions
  -h, --help               help for list-access
      --page int           Page number
      --per-page int       Number of items per page (default 25)
      --protocol string    Filter access log entries by protocol (one of: imap, pop3, smtp)
      --status string      Filter log entries by status (one of: Successful, Failed)
      --to-date string     Date range end (RFC 3339)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger mail logs](hostinger_mail_logs.md)	 - Logs commands

