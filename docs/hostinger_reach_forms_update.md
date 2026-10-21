## hostinger reach forms update

Update form

### Synopsis

Rename a form and/or change its status.

Setting a `draft` or `paused` form to `active` takes it live: its latest saved design is published
and visitors can submit it. `paused` and `draft` stop the form from accepting submissions. A form is
not reachable by visitors until it is taken live. Design changes saved while the form is already
`active` stay unpublished until the form is paused and set to `active` again. Renaming does not rename
the automatic `form:{name}` tag.

```
hostinger reach forms update <profile-uuid> <form-uuid> [flags]
```

### Options

```
  -h, --help            help for update
      --name string     Form name. Letters, digits, spaces and - . ' _ only.
      --status active   active takes the form live, `paused` and `draft` stop it accepting submissions. (one of: active, paused, draft)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach forms](hostinger_reach_forms.md)	 - Forms commands

