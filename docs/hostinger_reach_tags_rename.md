## hostinger reach tags rename

Rename a tag

### Synopsis

Rename a tag.

The contacts assigned to the tag are unaffected. Names are unique within a profile, so
renaming a tag to a name that is already taken is rejected.

```
hostinger reach tags rename <profile-uuid> <tag-uuid> [flags]
```

### Examples

```
  hostinger reach tags rename 550e8400-e09b-41d4-a716-400055000000 550e8400-e29b-41d4-a716-446655440000 --value Newsletter
```

### Options

```
  -h, --help           help for rename
      --value string   New tag name
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

* [hostinger reach tags](hostinger_reach_tags.md)	 - Tags commands

