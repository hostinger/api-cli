## hostinger reach tags assign-contacts-to

Assign contacts to a tag

### Synopsis

Assign a tag to many contacts at once.

Pass `contact_uuids` to target specific contacts, or `all_contacts` to target every contact
in the profile. The work is queued, so a success response means it was accepted rather than
finished. Contacts that already carry the tag are left alone.

```
hostinger reach tags assign-contacts-to <profile-uuid> <tag-uuid> [flags]
```

### Options

```
      --all-contacts            Apply to every contact in the profile
      --contact-uuids strings   Contacts to apply the change to. Required unless all_contacts is true.
  -h, --help                    help for assign-contacts-to
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach tags](hostinger_reach_tags.md)	 - Tags commands

