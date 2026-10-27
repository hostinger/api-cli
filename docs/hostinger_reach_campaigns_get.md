## hostinger reach campaigns get

Get campaign details

### Synopsis

Get a single campaign with its sender, subject, template reference, targeting and delivery
progress.

This describes how the campaign was set up and how far it has got. For opens, clicks and
unsubscribes use the campaign statistics endpoint. The email content is not included: pass
`template_uuid` to the get email template endpoint to read its HTML.

```
hostinger reach campaigns get <profile-uuid> <campaign-uuid> [flags]
```

### Examples

```
  hostinger reach campaigns get 550e8400-e09b-41d4-a716-400055000000 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
  -h, --help   help for get
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

* [hostinger reach campaigns](hostinger_reach_campaigns.md)	 - Campaigns commands

