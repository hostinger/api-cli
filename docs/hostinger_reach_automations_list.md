## hostinger reach automations list

List automations

### Synopsis

Get a paginated list of the automations in a profile.

Every automation comes with the counts of contacts that entered it, are moving through it,
finished it or failed on the way. Those counts describe the contact journey and are not
email engagement metrics - for opens, clicks and unsubscribes use the campaign statistics
endpoint instead.

```
hostinger reach automations list <profile-uuid> [flags]
```

### Examples

```
  hostinger reach automations list 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
      --all                  Fetch every page and print all items
  -h, --help                 help for list
      --page int             Page number
      --per-page int         Number of items per page (default 25)
      --sort-direction asc   Order automations by creation date. Newest first unless set to asc. (one of: asc, desc)
      --status completed     Filter automations by status. There is no completed status. An automation that has finished for every contact still reports `active`. (one of: active, paused, draft)
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

* [hostinger reach automations](hostinger_reach_automations.md)	 - Automations commands

