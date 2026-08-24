## hostinger reach segments preview-contacts-matching-conditions

Preview contacts matching conditions

### Synopsis

Preview the contacts matching a set of conditions without saving a segment.

The body is the same set of conditions accepted when creating or updating a segment, so this
is how to check who a filter reaches, and how many, before persisting it. Nothing is stored
and no contact is modified.

Call the segment filter attributes endpoint first to discover the valid `attribute`,
`operator` and `value` combinations.

```
hostinger reach segments preview-contacts-matching-conditions <profile-uuid> [flags]
```

### Options

```
      --conditions string       Conditions a contact must satisfy to appear in the preview (JSON)
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
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach segments](hostinger_reach_segments.md)	 - Segments commands

