## hostinger mail autoreplies create

Create autoreply

### Synopsis

Create an automatic reply for the given mailbox. A mailbox can have
only one autoreply. Omit `starts_at` to activate the autoreply
immediately and omit `ends_at` to keep it active indefinitely.

```
hostinger mail autoreplies create <mailbox-id> [flags]
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
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger mail autoreplies](hostinger_mail_autoreplies.md)	 - Autoreplies commands

