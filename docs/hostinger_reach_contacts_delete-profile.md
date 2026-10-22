## hostinger reach contacts delete-profile

Delete a profile contact

### Synopsis

Permanently delete a contact from a profile.

The contact is removed together with its custom field values and tag assignments.

```
hostinger reach contacts delete-profile <profile-uuid> <contact-uuid> [flags]
```

### Examples

```
  hostinger reach contacts delete-profile 550e8400-e09b-41d4-a716-400055000000 550e8400-e29b-41d4-a716-446655440000
```

### Options

```
  -h, --help   help for delete-profile
  -y, --yes    Skip the confirmation prompt
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

