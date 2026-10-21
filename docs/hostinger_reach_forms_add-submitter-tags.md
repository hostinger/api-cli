## hostinger reach forms add-submitter-tags

Add form submitter tags

### Synopsis

Add tags to a form so that every contact who submits it from now on receives them.

A tag the form already has is skipped; an unknown tag is rejected. Contacts who submitted earlier are
not tagged.

```
hostinger reach forms add-submitter-tags <profile-uuid> <form-uuid> [flags]
```

### Options

```
  -h, --help                help for add-submitter-tags
      --tag-uuids strings   
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach forms](hostinger_reach_forms.md)	 - Forms commands

