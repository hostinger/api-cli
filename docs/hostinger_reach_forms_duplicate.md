## hostinger reach forms duplicate

Duplicate form

### Synopsis

Copy a form into a new draft form with the same name and design.

The copy shares the source's automatic `form:{name}` tag; other submitter tags are not copied. The
copy is not live until its status is set to `active`.

```
hostinger reach forms duplicate <profile-uuid> <form-uuid> [flags]
```

### Options

```
  -h, --help   help for duplicate
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach forms](hostinger_reach_forms.md)	 - Forms commands

