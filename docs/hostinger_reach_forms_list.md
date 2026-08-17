## hostinger reach forms list

List forms

### Synopsis

Get a paginated list of the signup forms in a profile.

Each form carries a reference to the template that renders it. Get the form details for a
directly usable template URL and for the tags the form puts on the contacts it captures.

```
hostinger reach forms list <profile-uuid> [flags]
```

### Options

```
  -h, --help           help for list
      --page int       Page number
      --per-page int   Number of items per page (default 25)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach forms](hostinger_reach_forms.md)	 - Forms commands

