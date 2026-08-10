## hostinger reach tags remove-contacts-from

Remove contacts from a tag

### Synopsis

Remove a tag from many contacts at once.

Pass `contact_uuids` to target specific contacts, or `all_contacts` to target every contact
in the profile. The work is queued, so a success response means it was accepted rather than
finished. The tag itself and the contacts are not deleted.

```
hostinger reach tags remove-contacts-from <profile-uuid> <tag-uuid> [flags]
```

### Options

```
      --all-contacts            Apply to every contact in the profile
      --contact-uuids strings   Contacts to apply the change to. Required unless all_contacts is true.
  -h, --help                    help for remove-contacts-from
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach tags](hostinger_reach_tags.md)	 - Tags commands

