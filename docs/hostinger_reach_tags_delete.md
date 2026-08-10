## hostinger reach tags delete

Delete a tag

### Synopsis

Delete a tag and remove it from every contact carrying it.

The contacts themselves are not deleted. This is idempotent: deleting a tag that does not
exist in the profile still succeeds.

```
hostinger reach tags delete <profile-uuid> <tag-uuid> [flags]
```

### Options

```
  -h, --help   help for delete
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach tags](hostinger_reach_tags.md)	 - Tags commands

