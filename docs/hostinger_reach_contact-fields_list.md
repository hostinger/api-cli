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

### Options

```
  -h, --help   help for list
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach contact-fields](hostinger_reach_contact-fields.md)	 - Contact Fields commands

