## hostinger reach contacts list

List contacts

### Synopsis

Get a list of contacts, optionally filtered by group and subscription status.

This endpoint returns a paginated list of contacts with their basic information.
You can filter contacts by group UUID and subscription status.

Deprecated. This endpoint cannot target a profile, so it always falls back to the
client's default profile and cannot list contacts of any other profile. Use
`GET /api/reach/v1/profiles/{profileUuid}/contacts` instead, which also replaces the
group filter with a tag filter.

```
hostinger reach contacts list [flags]
```

### Examples

```
  hostinger reach contacts list
```

### Options

```
      --all                          Fetch every page and print all items
      --group-uuid string            Filter contacts by group UUID
  -h, --help                         help for list
      --page int                     Page number
      --subscription-status string   Filter contacts by subscription status (one of: subscribed, unsubscribed, confirmed, pending)
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --domain string     Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger reach contacts](hostinger_reach_contacts.md)	 - Contacts commands

