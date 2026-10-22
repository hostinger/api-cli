## hostinger mail forwarders update-keep-copy-setting

Update forwarder keep-copy setting

### Synopsis

Enable or disable keeping a copy of forwarded messages in the
mailbox.

```
hostinger mail forwarders update-keep-copy-setting <forwarder-id> [flags]
```

### Examples

```
  hostinger mail forwarders update-keep-copy-setting FW1a2b3c4d5e6f7g --is-keep-copy-enabled=true
```

### Options

```
  -h, --help                   help for update-keep-copy-setting
      --is-keep-copy-enabled   Whether to keep a copy of forwarded messages in the mailbox
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

* [hostinger mail forwarders](hostinger_mail_forwarders.md)	 - Forwarders commands

