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

### Examples

```
  hostinger reach forms create 550e8400-e09b-41d4-a716-400055000000 --name 'Newsletter signup'
```

### Options

```
  -h, --help          help for create
      --name string   Form name. Letters, digits, spaces and - . ' _ only.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger reach forms](hostinger_reach_forms.md)	 - Forms commands

