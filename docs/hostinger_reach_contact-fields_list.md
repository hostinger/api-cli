## hostinger reach contact-fields list

List contact fields

### Synopsis

Get the custom contact fields defined in a profile.

Custom fields let you store your own attributes on contacts. The returned uuids are what
you pass to the contact update endpoint to set values, and choice fields also list the
options available to pick from.

```
hostinger reach contact-fields list <profile-uuid> [flags]
```

### Examples

```
  hostinger reach contact-fields list 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
  -h, --help   help for list
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

* [hostinger reach contact-fields](hostinger_reach_contact-fields.md)	 - Contact Fields commands

