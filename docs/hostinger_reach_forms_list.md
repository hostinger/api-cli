## hostinger reach forms list

List forms

### Synopsis

Get a paginated list of the signup forms in a profile.

Each form carries a reference to the template that renders it. Get the form details for a
directly usable template URL and for the tags the form puts on the contacts it captures.

```
hostinger reach forms list <profile-uuid> [flags]
```

### Examples

```
  hostinger reach forms list 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
      --all             Fetch every page and print all items
  -h, --help            help for list
      --page int        Page number
      --per-page int    Number of items per page (default 25)
      --status string   Return only forms with this status. (one of: active, paused, draft)
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

* [hostinger reach forms](hostinger_reach_forms.md)	 - Forms commands

