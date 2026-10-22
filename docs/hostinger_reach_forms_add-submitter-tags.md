## hostinger reach forms add-submitter-tags

Add form submitter tags

### Synopsis

Add tags to a form so that every contact who submits it from now on receives them.

A tag the form already has is skipped; an unknown tag is rejected. Contacts who submitted earlier are
not tagged.

```
hostinger reach forms add-submitter-tags <profile-uuid> <form-uuid> [flags]
```

### Examples

```
  hostinger reach forms add-submitter-tags 550e8400-e09b-41d4-a716-400055000000 550e8400-e09b-41d4-a716-400055000000 --tag-uuids 550e8400-e29b-41d4-a716-446655440000
```

### Options

```
  -h, --help                help for add-submitter-tags
      --tag-uuids strings   
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

