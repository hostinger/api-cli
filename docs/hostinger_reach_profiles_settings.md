## hostinger reach profiles settings

Get profile settings

### Synopsis

Get the profile's saved sender identity and brand details.

Read the default sender before creating a campaign instead of asking the customer, and
reuse the brand fields when generating email content. Any field can be `null` for a profile
that never filled its settings in. This does not check that the sending domain is
verified, use the connected sending domain endpoint for that.

```
hostinger reach profiles settings <profile-uuid> [flags]
```

### Examples

```
  hostinger reach profiles settings 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
  -h, --help   help for settings
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

* [hostinger reach profiles](hostinger_reach_profiles.md)	 - Profiles commands

