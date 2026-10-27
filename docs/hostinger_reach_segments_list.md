## hostinger reach segments list

List segments

### Synopsis

Get a list of all contact segments.

This endpoint returns a list of contact segments that can be used to organize contacts.

Deprecated. This endpoint cannot target a profile, so it always falls back to
the client's default profile and cannot list the segments of any other profile. Use
`GET /api/reach/v1/profiles/{profileUuid}/segmentation/segments` instead.

```
hostinger reach segments list [flags]
```

### Examples

```
  hostinger reach segments list
```

### Options

```
  -h, --help   help for list
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

