## hostinger reach tags rename

Rename a tag

### Synopsis

Rename a tag.

The contacts assigned to the tag are unaffected. Names are unique within a profile, so
renaming a tag to a name that is already taken is rejected.

```
hostinger reach tags rename <profile-uuid> <tag-uuid> [flags]
```

### Options

```
  -h, --help           help for rename
      --value string   New tag name
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach tags](hostinger_reach_tags.md)	 - Tags commands

