## hostinger reach forms delete

Delete form

### Synopsis

Permanently delete a form together with its template.

A form that has already captured submissions cannot be deleted, so that the contacts it collected
are never silently discarded - pause the form instead to stop it collecting new ones. Views alone
do not block deletion.

```
hostinger reach forms delete <profile-uuid> <form-uuid> [flags]
```

### Examples

```
  hostinger reach forms delete 550e8400-e09b-41d4-a716-400055000000 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
  -h, --help   help for delete
  -y, --yes    Skip the confirmation prompt
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

