## hostinger reach segments create-profile

Create a profile segment

### Synopsis

Create a segment in a profile.

A segment is a saved set of conditions rather than a fixed list, so its membership changes
as contacts change. Creating one does not modify any contact.

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
hostinger reach segments create-profile <profile-uuid> [flags]
```

### Examples

```
  hostinger reach segments create-profile 550e8400-e09b-41d4-a716-400055000000 --conditions @conditions.json --logic AND --name 'VIP Customers'
```

### Options

```
      --conditions string   Conditions a contact must satisfy to fall into the segment (JSON, @file or -; shape in the help above)
  -h, --help                help for create-profile
      --logic string        How to combine multiple conditions (one of: AND, OR)
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

