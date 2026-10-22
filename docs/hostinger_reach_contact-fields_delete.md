## hostinger reach contact-fields delete

Delete a contact field

### Synopsis

Delete a custom contact field.

Every value contacts hold for the field is deleted with it, and for the choice types so
are its options. The contacts themselves are not affected.

```
hostinger reach contact-fields delete <profile-uuid> <field-uuid> [flags]
```

### Examples

```
  hostinger reach contact-fields delete 550e8400-e09b-41d4-a716-400055000000 550e8400-e29b-41d4-a716-446655440000
```

### Options

```
  -h, --help   help for delete
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

* [hostinger reach contact-fields](hostinger_reach_contact-fields.md)	 - Contact Fields commands

