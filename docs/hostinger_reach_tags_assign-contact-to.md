## hostinger reach tags assign-contact-to

Assign a contact to a tag

### Synopsis

Assign a tag to a single contact.

Unlike the bulk endpoint this is applied immediately rather than queued. Assigning a tag
the contact already carries succeeds without duplicating it.

```
hostinger reach tags assign-contact-to <profile-uuid> <tag-uuid> <contact-uuid> [flags]
```

### Options

```
  -h, --help   help for assign-contact-to
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach tags](hostinger_reach_tags.md)	 - Tags commands

