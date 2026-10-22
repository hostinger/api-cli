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

### Examples

```
  hostinger reach tags create-or-find 550e8400-e09b-41d4-a716-400055000000 --names Newsletter,VIP,Promo
```

### Options

```
  -h, --help            help for create-or-find
      --names strings   
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

