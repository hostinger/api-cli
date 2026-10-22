## hostinger hosting cron-jobs list

List account cron jobs

### Synopsis

Returns the list of cron jobs configured for the specified account, including their schedule and command.

```
hostinger hosting cron-jobs list <username> [flags]
```

### Examples

```
  hostinger hosting cron-jobs list u123456789
```

### Options

```
  -h, --help   help for list
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

* [hostinger hosting cron-jobs](hostinger_hosting_cron-jobs.md)	 - Cron Jobs commands

