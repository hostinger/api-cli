## hostinger reach forms delete

Delete form

### Synopsis

Permanently delete a form together with its template.

A form that has already captured submissions cannot be deleted, so that the contacts it collected
are never silently discarded - pause the form instead to stop it collecting new ones. Views alone
do not block deletion.

```
hostinger reach forms delete <profile-uuid> <form-uuid> [flags]
```

### Options

```
  -h, --help   help for delete
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach forms](hostinger_reach_forms.md)	 - Forms commands

