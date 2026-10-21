## hostinger reach forms remove-submitter-tag

Remove form submitter tag

### Synopsis

Stop a form from applying a tag to its future submitters.

Succeeds also when the form did not have the tag. Contacts who already received the tag keep it, and
the tag itself is not deleted. The automatic `form:{name}` tag cannot be removed and is rejected with
a conflict.

```
hostinger reach forms remove-submitter-tag <profile-uuid> <form-uuid> [flags]
```

### Options

```
  -h, --help              help for remove-submitter-tag
      --tag-uuid string   
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach forms](hostinger_reach_forms.md)	 - Forms commands

