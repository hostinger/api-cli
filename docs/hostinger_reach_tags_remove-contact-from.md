## hostinger reach tags remove-contact-from

Remove a contact from a tag

### Synopsis

Remove a tag from a single contact.

Unlike the bulk endpoint this is applied immediately rather than queued. Neither the tag
nor the contact is deleted.

```
hostinger reach tags remove-contact-from <profile-uuid> <tag-uuid> <contact-uuid> [flags]
```

### Options

```
  -h, --help   help for remove-contact-from
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach tags](hostinger_reach_tags.md)	 - Tags commands

