## hostinger reach segments preview-contacts-matching-conditions

Preview contacts matching conditions

### Synopsis

Preview the contacts matching a set of conditions without saving a segment.

The body is the same set of conditions accepted when creating or updating a segment, so this
is how to check who a filter reaches, and how many, before persisting it. Nothing is stored
and no contact is modified.

Call the segment filter attributes endpoint first to discover the valid `attribute`,
`operator` and `value` combinations.

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
hostinger reach segments preview-contacts-matching-conditions <profile-uuid> [flags]
```

### Examples

```
  hostinger reach segments preview-contacts-matching-conditions 550e8400-e09b-41d4-a716-400055000000 --conditions @conditions.json --logic AND
```

### Options

```
      --conditions string       Conditions a contact must satisfy to appear in the preview (JSON, @file or -; shape in the help above)
  -h, --help                    help for preview-contacts-matching-conditions
      --logic string            How to combine multiple conditions (one of: AND, OR)
      --page int                Page number
      --per-page int            Number of items per page
      --search string           Narrow the preview to contacts whose email matches
      --sort-by string          (one of: email, name, surname, phone, subscription_status)
      --sort-direction string   (one of: asc, desc)
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

