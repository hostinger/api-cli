## hostinger reach segments list-contacts

List segment contacts

### Synopsis

Retrieve contacts associated with a specific segment.

This endpoint allows you to fetch and filter contacts that belong to a particular segment,
identified by its UUID.

Deprecated. This endpoint cannot target a profile, so it always falls back to
the client's default profile and cannot read segments of any other profile. Use
`GET /api/reach/v1/profiles/{profileUuid}/segmentation/segments/{segmentUuid}/contacts` instead.

```
hostinger reach segments list-contacts <segment-uuid> [flags]
```

### Examples

```
  hostinger reach segments list-contacts 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
      --all            Fetch every page and print all items
  -h, --help           help for list-contacts
      --page int       Page number
      --per-page int   Number of items per page (default 25)
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

* [hostinger reach segments](hostinger_reach_segments.md)	 - Segments commands

