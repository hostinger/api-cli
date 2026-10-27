## hostinger reach segments get

Get segment details

### Synopsis

Get details of a specific segment.

This endpoint retrieves information about a single segment identified by UUID.
Segments are used to organize and group contacts based on specific criteria.

Deprecated. This endpoint cannot target a profile, so it always falls back to
the client's default profile and cannot read segments of any other profile. Use
`GET /api/reach/v1/profiles/{profileUuid}/segmentation/segments/{segmentUuid}` instead.

```
hostinger reach segments get <segment-uuid> [flags]
```

### Examples

```
  hostinger reach segments get 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
  -h, --help   help for get
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

