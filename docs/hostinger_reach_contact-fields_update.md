## hostinger reach contact-fields update

Update a contact field

### Synopsis

Rename a custom contact field and, for the choice types, replace its option set.

Options carrying a uuid are kept and relabelled, options without one are created, and any
existing option left out of the list is deleted along with the values contacts hold for
it. The field type and slug cannot be changed.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --options
    [
      {
        "label": "string",
        "uuid": "string"
      }
    ]

```
hostinger reach contact-fields update <profile-uuid> <field-uuid> [flags]
```

### Examples

```
  hostinger reach contact-fields update 550e8400-e09b-41d4-a716-400055000000 550e8400-e29b-41d4-a716-446655440000 --label 'Job title'
```

### Options

```
  -h, --help             help for update
      --label string     
      --options string   Replaces the option set when provided. Entries carrying a uuid are kept and relabelled, entries without one are created, and any existing option missing from the list is deleted along with the values contacts hold for it. (JSON, @file or -; shape in the help above)
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

