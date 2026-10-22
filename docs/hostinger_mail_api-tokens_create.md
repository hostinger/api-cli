## hostinger mail api-tokens create

Create API token

### Synopsis

Create an API token for the given mail order. The token grants access
to the Hostinger Email API (https://api.mail.hostinger.com/), where
you can provision and manage the mailboxes it is scoped to.

The plaintext token is returned only in this response, never again.
A maximum of 10 tokens can exist per order. Use
`scope.has_all_mailboxes` to cover all current and future mailboxes,
or list specific mailboxes in `scope.mailbox_ids`.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --scope
    {
      "has_all_mailboxes": false,
      "mailbox_ids": [
        "AC1a2b3c4d5e6f7g"
      ]
    }

```
hostinger mail api-tokens create <order-id> [flags]
```

### Examples

```
  hostinger mail api-tokens create OR1a2b3c4d5e6f7g --name 'CRM integration' --scope @scope.json
```

### Options

```
  -h, --help           help for create
      --name string    Human-readable label for this token
      --scope string   Mailbox scope this token can access (JSON, @file or -; shape in the help above)
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

* [hostinger mail api-tokens](hostinger_mail_api-tokens.md)	 - API Tokens commands

