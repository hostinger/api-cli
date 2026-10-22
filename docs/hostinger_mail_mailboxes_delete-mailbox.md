## hostinger mail mailboxes delete-mailbox

Delete mailbox

### Synopsis

Delete a mailbox. The mailbox is soft-deleted and stays restorable
for a limited period before it is permanently removed.

```
hostinger mail mailboxes delete-mailbox <mailbox-id> [flags]
```

### Examples

```
  hostinger mail mailboxes delete-mailbox AC1a2b3c4d5e6f7g
```

### Options

```
  -h, --help   help for delete-mailbox
  -y, --yes    Skip the confirmation prompt
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

