## hostinger reach tags create-or-find

Create or find tags

### Synopsis

Create tags in a profile.

Names that already exist in the profile are not duplicated: the existing tag is returned
instead, so the call is safe to repeat. Every tag in the request is returned, whether it
was created now or already existed.

```
hostinger reach tags create-or-find <profile-uuid> [flags]
```

### Options

```
  -h, --help            help for create-or-find
      --names strings   
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach tags](hostinger_reach_tags.md)	 - Tags commands

