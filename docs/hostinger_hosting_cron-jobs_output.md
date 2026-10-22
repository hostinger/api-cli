## hostinger hosting cron-jobs output

Get cron job output

### Synopsis

Returns the output captured from the last execution of the cron job identified by its uid.

The uid is returned by the list cron jobs endpoint.

```
hostinger hosting cron-jobs output <username> <uid> [flags]
```

### Examples

```
  hostinger hosting cron-jobs output u123456789 cron_abc123
```

### Options

```
  -h, --help   help for output
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

