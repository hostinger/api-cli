## hostinger mail autoreplies create

Create autoreply

### Synopsis

Create an automatic reply for the given mailbox. A mailbox can have
only one autoreply. Omit `starts_at` to activate the autoreply
immediately and omit `ends_at` to keep it active indefinitely.

```
hostinger mail autoreplies create <mailbox-id> [flags]
```

### Examples

```
  hostinger mail autoreplies create AC1a2b3c4d5e6f7g --body 'I am on vacation until August 1st.' --subject 'Out of office'
```

### Options

```
      --body string           Body of the automatic reply
      --display-name string   Sender display name used for the reply
      --ends-at string        When the autoreply stops. Omit for an indefinite autoreply.
  -h, --help                  help for create
      --starts-at string      When the autoreply becomes active. Defaults to now.
      --subject string        Subject of the automatic reply
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

* [hostinger mail autoreplies](hostinger_mail_autoreplies.md)	 - Autoreplies commands

