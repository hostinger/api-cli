## hostinger mail mailboxes mailbox

Get mailbox list

### Synopsis

Retrieve a paginated list of mailboxes belonging to a mail order.

Use this endpoint to monitor mailboxes of your mail service, including
their status, enabled protocols, attached resource counts, and
periodically synced usage numbers (usage may lag behind live values).

```
hostinger mail mailboxes mailbox <order-id> [flags]
```

### Options

```
  -h, --help            help for mailbox
      --page int        Page number
      --per-page int    Number of items per page (default 25)
      --search string   Filter mailboxes whose email address contains the given string
      --sort -          Sort mailboxes by field. Prefix with - for descending order. (one of: address, -address) (default "address")
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger mail mailboxes](hostinger_mail_mailboxes.md)	 - Mailboxes commands

