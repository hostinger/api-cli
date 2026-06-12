## hapi reach contacts list

List contacts

### Synopsis

Get a list of contacts, optionally filtered by group and subscription status.

This endpoint returns a paginated list of contacts with their basic information.
You can filter contacts by group UUID and subscription status.

```
hapi reach contacts list [flags]
```

### Options

```
      --group-uuid string            Filter contacts by group UUID
  -h, --help                         help for list
      --page int                     Page number
      --subscription-status string   Filter contacts by subscription status (one of: subscribed, unsubscribed)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi reach contacts](hapi_reach_contacts.md)	 - Contacts commands

