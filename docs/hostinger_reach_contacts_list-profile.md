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

### Options

```
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
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach contacts](hostinger_reach_contacts.md)	 - Contacts commands

