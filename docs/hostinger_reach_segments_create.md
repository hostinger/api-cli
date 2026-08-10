## hostinger reach segments create

Create a new contact segment

### Synopsis

Create a new contact segment.

This endpoint allows creating a new contact segment that can be used to organize contacts.
The segment can be configured with specific criteria like email, name, subscription status, etc.

**Deprecated.** This endpoint cannot target a profile, so it always falls back to
the client's default profile and cannot create segments in any other profile. Use
`POST /api/reach/v1/profiles/{profileUuid}/segmentation/segments` instead.

```
hostinger reach segments create [flags]
```

### Options

```
      --conditions string    (JSON)
  -h, --help                help for create
      --logic string        (one of: AND, OR)
      --name string         
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach segments](hostinger_reach_segments.md)	 - Segments commands

