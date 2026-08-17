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

### Options

```
  -h, --help                 help for list
      --page int             Page number
      --per-page int         Number of items per page (default 25)
      --sort-direction asc   Order automations by creation date. Newest first unless set to asc. (one of: asc, desc)
      --status completed     Filter automations by status.
                             
                             There is no completed status. An automation that has finished for every contact still
                             reports `active`. (one of: active, paused, draft)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach automations](hostinger_reach_automations.md)	 - Automations commands

