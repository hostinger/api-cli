## hostinger mail autoreplies update

Update autoreply

### Synopsis

Replace the autoreply with the given content and schedule. Omitted
optional fields are cleared: omit `starts_at` to activate the
autoreply immediately and omit `ends_at` to keep it active
indefinitely.

```
hostinger mail autoreplies update <autoreply-id> [flags]
```

### Options

```
      --body string           Body of the automatic reply
      --display-name string   Sender display name used for the reply
      --ends-at string        When the autoreply stops. Omit for an indefinite autoreply.
  -h, --help                  help for update
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

