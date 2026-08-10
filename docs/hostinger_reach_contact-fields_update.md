## hostinger reach contact-fields update

Update a contact field

### Synopsis

Rename a custom contact field and, for the choice types, replace its option set.

Options carrying a uuid are kept and relabelled, options without one are created, and any
existing option left out of the list is deleted along with the values contacts hold for
it. The field type and slug cannot be changed.

```
hostinger reach contact-fields update <profile-uuid> <field-uuid> [flags]
```

### Options

```
  -h, --help             help for update
      --label string     
      --options string   Replaces the option set when provided. Entries carrying a uuid are kept and relabelled, entries without one are created, and any existing option missing from the list is deleted along with the values contacts hold for it. (JSON)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach contact-fields](hostinger_reach_contact-fields.md)	 - Contact Fields commands

