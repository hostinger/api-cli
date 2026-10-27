## hostinger reach segments update-profile

Update a profile segment

### Synopsis

Rename a segment and/or replace the conditions that define it.

`name` is always required. Omit `conditions` to rename without touching the conditions;
supply them and they replace the existing set entirely rather than being merged into it.
Contacts are never modified, but which of them match the segment can change immediately.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --conditions
    [
      {
        "attribute": "email",
        "operator": "equals",
        "value": "example@example.com"
      }
    ]

```
hostinger reach segments update-profile <profile-uuid> <segment-uuid> [flags]
```

### Examples

```
  hostinger reach segments update-profile 550e8400-e09b-41d4-a716-400055000000 550e8400-e09b-41d4-a716-400055000000 --name 'VIP Customers'
```

### Options

```
      --conditions string   Replaces the existing conditions entirely. Omit to keep the current ones. (JSON, @file or -; shape in the help above)
  -h, --help                help for update-profile
      --logic string        How to combine multiple conditions. Required when conditions are given. (one of: AND, OR)
      --name string         
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

