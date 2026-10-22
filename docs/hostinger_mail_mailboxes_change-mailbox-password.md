## hostinger mail mailboxes change-mailbox-password

Change mailbox password

### Synopsis

Change the password of a mailbox.

```
hostinger mail mailboxes change-mailbox-password <mailbox-id> [flags]
```

### Examples

```
  hostinger mail mailboxes change-mailbox-password AC1a2b3c4d5e6f7g --password 'SecurePassword123!'
```

### Options

```
  -h, --help              help for change-mailbox-password
      --password string   New mailbox password. Minimum 8 characters with uppercase, lowercase, number and special character; must not be a commonly used password.
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

* [hostinger mail mailboxes](hostinger_mail_mailboxes.md)	 - Mailboxes commands

