## hostinger mail mailboxes list

List mailboxes

### Synopsis

Retrieve a paginated list of mailboxes belonging to a mail order.

Use this endpoint to monitor mailboxes of your mail service, including
their status, enabled protocols, attached resource counts, and
periodically synced usage numbers (usage may lag behind live values).

The order can be given as [order-id] or left out, in which case it comes from --order-id, HOSTINGER_ORDER_ID, the .hostinger/site.json written by `hostinger up` or `hostinger link` (hosting orders), or a prompt when run in a terminal (the only order is used without asking).

```
hostinger mail mailboxes list [order-id] [flags]
```

### Examples

```
  hostinger mail mailboxes list OR1a2b3c4d5e6f7g
```

### Options

```
      --all             Fetch every page and print all items
  -h, --help            help for list
      --page int        Page number
      --per-page int    Number of items per page (default 25)
      --search string   Filter mailboxes whose email address contains the given string
      --sort -          Sort mailboxes by field. Prefix with - for descending order. (one of: address, -address) (default "address")
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

* [hostinger mail mailboxes](hostinger_mail_mailboxes.md)	 - Mailboxes commands

