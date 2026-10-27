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

### Examples

```
  hostinger mail autoreplies update AR1a2b3c4d5e6f7g --body 'I am on vacation until August 1st.' --subject 'Out of office'
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

* [hostinger mail autoreplies](hostinger_mail_autoreplies.md)	 - Autoreplies commands

