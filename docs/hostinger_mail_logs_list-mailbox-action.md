## hostinger mail logs list-mailbox-action

List mailbox action logs

### Synopsis

Retrieve paginated mailbox action logs (message and mailbox events)
for a mailbox in the given mail order. The mailbox email must belong
to the order's domain. Supports date range and event type filters.
Results are sorted by timestamp descending.

The order can be given as [order-id] or left out, in which case it comes from --order-id, HOSTINGER_ORDER_ID, the .hostinger/site.json written by `hostinger up` or `hostinger link` (hosting orders), or a prompt when run in a terminal (the only order is used without asking).

```
hostinger mail logs list-mailbox-action [order-id] [flags]
```

### Examples

```
  hostinger mail logs list-mailbox-action OR1a2b3c4d5e6f7g --email user@example.com
```

### Options

```
      --all                Fetch every page and print all items
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
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --domain string     Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger mail logs](hostinger_mail_logs.md)	 - Logs commands

