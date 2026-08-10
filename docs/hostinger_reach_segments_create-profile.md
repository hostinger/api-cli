## hostinger reach segments create-profile

Create a profile segment

### Synopsis

Create a segment in a profile.

A segment is a saved set of conditions rather than a fixed list, so its membership changes
as contacts change. Creating one does not modify any contact.

```
hostinger reach segments create-profile <profile-uuid> [flags]
```

### Options

```
      --conditions string   Conditions a contact must satisfy to fall into the segment (JSON)
  -h, --help                help for create-profile
      --logic string        How to combine multiple conditions (one of: AND, OR)
      --name string         
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach segments](hostinger_reach_segments.md)	 - Segments commands

