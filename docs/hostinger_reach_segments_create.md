## hostinger reach segments create

Create a new contact segment

### Synopsis

Create a new contact segment.

This endpoint allows creating a new contact segment that can be used to organize contacts.
The segment can be configured with specific criteria like email, name, subscription status, etc.

Deprecated. This endpoint cannot target a profile, so it always falls back to
the client's default profile and cannot create segments in any other profile. Use
`POST /api/reach/v1/profiles/{profileUuid}/segmentation/segments` instead.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --conditions
    [
      {
        "attribute": "note",
        "operator": "equals",
        "value": "string"
      }
    ]

```
hostinger reach segments create [flags]
```

### Examples

```
  hostinger reach segments create --conditions '[{"attribute":"note","operator":"equals","value":"string"}]' --logic AND --name 'New segment name'
```

### Options

```
      --conditions string   (JSON, @file or -; shape in the help above)
  -h, --help                help for create
      --logic string        (one of: AND, OR)
      --name string         
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

