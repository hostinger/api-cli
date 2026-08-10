## hostinger reach segments list

List segments

### Synopsis

Get a list of all contact segments.

This endpoint returns a list of contact segments that can be used to organize contacts.

**Deprecated.** This endpoint cannot target a profile, so it always falls back to
the client's default profile and cannot list the segments of any other profile. Use
`GET /api/reach/v1/profiles/{profileUuid}/segmentation/segments` instead.

```
hostinger reach segments list [flags]
```

### Options

```
  -h, --help   help for list
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach segments](hostinger_reach_segments.md)	 - Segments commands

