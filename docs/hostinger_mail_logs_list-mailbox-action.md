## hostinger mail logs list-mailbox-action

List mailbox action logs

### Synopsis

Retrieve paginated mailbox action logs (message and mailbox events)
for a mailbox in the given mail order. The mailbox email must belong
to the order's domain. Supports date range and event type filters.
Results are sorted by timestamp descending.

```
hostinger mail logs list-mailbox-action <order-id> [flags]
```

### Options

```
      --date from_date     Exact date filter (YYYY-MM-DD). Takes precedence over from_date/`to_date` when both are given.
      --email string       Mailbox email address. Must belong to the order's domain.
      --event string       Filter mailbox action log entries by event type (one of: MessageNew, MessageRead, MessageAppend, MessageExpunge, MailboxCreate, MailboxDelete, MailboxRename)
      --from-date string   Date range start (RFC 3339)
  -h, --help               help for list-mailbox-action
      --page int           Page number
      --per-page int       Number of items per page (default 25)
      --to-date string     Date range end (RFC 3339)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger mail logs](hostinger_mail_logs.md)	 - Logs commands

