## hostinger reach contacts get

Get contact details

### Synopsis

Get the full details of a single contact.

Alongside the contact's own attributes this returns the tags assigned to it and the
values it holds for the profile's custom contact fields.

```
hostinger reach contacts get <profile-uuid> <contact-uuid> [flags]
```

### Examples

```
  hostinger reach contacts get 550e8400-e09b-41d4-a716-400055000000 550e8400-e29b-41d4-a716-446655440000
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

* [hostinger reach contacts](hostinger_reach_contacts.md)	 - Contacts commands

