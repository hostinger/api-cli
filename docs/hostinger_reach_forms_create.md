## hostinger reach forms create

Create form

### Synopsis

Create an empty draft form.

Reach also adds an automatic `form:{name}` tag that every submitter receives. A draft has no design
and is not reachable by visitors: save a design with the update form template operation, then take
the form live by setting its status to `active`. Creating a form does not consume AI credits.

```
hostinger reach forms create <profile-uuid> [flags]
```

### Options

```
  -h, --help          help for create
      --name string   Form name. Letters, digits, spaces and - . ' _ only.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach forms](hostinger_reach_forms.md)	 - Forms commands

