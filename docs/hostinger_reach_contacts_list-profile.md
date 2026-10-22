## hostinger reach contacts list-profile

List profile contacts

### Synopsis

Get a paginated list of contacts belonging to a profile.

Contacts can be filtered by subscription status, by tag, and by an email search term.
The `meta.total` field of the response is the number of contacts matching the filters,
so calling this endpoint without filters gives the profile's total contact count.

```
hostinger reach contacts list-profile <profile-uuid> [flags]
```

### Examples

```
  hostinger reach contacts list-profile 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
      --all                          Fetch every page and print all items
  -h, --help                         help for list-profile
      --page int                     Page number
      --per-page int                 Number of items per page (default 25)
      --search string                Search contacts by email
      --subscription-status string   Filter contacts by subscription status (one of: subscribed, unsubscribed, confirmed, pending)
      --tag-uuid string              Filter contacts by tag UUID
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

* [hostinger reach contacts](hostinger_reach_contacts.md)	 - Contacts commands

