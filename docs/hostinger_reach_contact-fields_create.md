## hostinger reach contact-fields create

Create a contact field

### Synopsis

Define a new custom contact field in a profile.

The `slug` is derived from the label and, like the field type, cannot be changed later.
Use the returned uuid to set values on contacts.

```
hostinger reach contact-fields create <profile-uuid> [flags]
```

### Options

```
  -h, --help              help for create
      --label string      
      --options strings   Required for single_choice and multi_choice, ignored for the scalar types. Labels must be unique regardless of casing.
      --type string       Immutable once the field exists (one of: text, number, date, single_choice, multi_choice)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach contact-fields](hostinger_reach_contact-fields.md)	 - Contact Fields commands

