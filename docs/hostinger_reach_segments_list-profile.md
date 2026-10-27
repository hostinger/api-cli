## hostinger reach segments list-profile

List profile segments

### Synopsis

Get a paginated list of the segments defined in a profile.

Each entry carries the number of contacts currently matching it, which is recalculated on
read rather than stored. Use `count_type` to count either every matching contact or only
the subscribed ones.

```
hostinger reach segments list-profile <profile-uuid> [flags]
```

### Examples

```
  hostinger reach segments list-profile 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
      --all                 Fetch every page and print all items
      --count-type string   Which matching contacts to count for each segment (one of: all, subscribed) (default "all")
  -h, --help                help for list-profile
      --page int            Page number
      --per-page int        Number of items per page (default 25)
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

* [hostinger reach segments](hostinger_reach_segments.md)	 - Segments commands

