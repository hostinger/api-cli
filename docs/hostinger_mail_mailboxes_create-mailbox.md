## hostinger mail mailboxes create-mailbox

Create mailbox

### Synopsis

Create a mailbox under the given mail order. The full email address is
composed from the given local part and the domain of the order.

The order can be given as [order-id] or left out, in which case it comes from --order-id, HOSTINGER_ORDER_ID, the .hostinger/site.json written by `hostinger up` or `hostinger link` (hosting orders), or a prompt when run in a terminal (the only order is used without asking).

```
hostinger mail mailboxes create-mailbox [order-id] [flags]
```

### Examples

```
  hostinger mail mailboxes create-mailbox OR1a2b3c4d5e6f7g --local-part john.doe --password 'SecurePassword123!'
```

### Options

```
  -h, --help                help for create-mailbox
      --local-part string   Local part of the mailbox address (the part before the @). The domain is taken from the order. Must start and end with a letter or digit; single dots, underscores and hyphens are allowed in between.
      --password string     Mailbox password. Minimum 8 characters with uppercase, lowercase, number and special character.
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

