## hostinger mail logs list-outbound

List outbound logs

### Synopsis

Retrieve paginated outbound (sent mail) delivery logs for the domain
attached to the given mail order. Supports filtering by account, date
range, status, sender, and recipient. Results are sorted by timestamp
descending.

```
hostinger mail logs list-outbound <order-id> [flags]
```

### Examples

```
  hostinger mail logs list-outbound OR1a2b3c4d5e6f7g
```

### Options

```
      --account string     Filter log entries by a specific email account
      --all                Fetch every page and print all items
      --date from_date     Exact date filter (YYYY-MM-DD). Takes precedence over from_date/`to_date` when both are given.
      --from-date string   Date range start (RFC 3339)
  -h, --help               help for list-outbound
      --page int           Page number
      --per-page int       Number of items per page (default 25)
      --recipient string   Filter log entries by recipient. Accepts a full email address or a domain.
      --sender string      Filter log entries by sender. Accepts a full email address or a domain.
      --status string      Filter log entries by status (one of: Successful, Failed)
      --to-date string     Date range end (RFC 3339)
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

* [hostinger mail logs](hostinger_mail_logs.md)	 - Logs commands

