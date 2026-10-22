## hostinger hosting nodejs build-logs

Get NodeJS build logs

### Synopsis

Retrieve logs from a specific Node.js build process.

To stream live output while a build is running, poll this endpoint repeatedly
while the build state is `running`, passing the previously returned `lines` count
as `from_line` to fetch only new output since the last call.
Log content may contain ANSI escape sequences (color codes).

```
hostinger hosting nodejs build-logs <username> <domain> <uuid> [flags]
```

### Examples

```
  hostinger hosting nodejs build-logs u123456789 mydomain.tld 123e4567-e89b-12d3-a456-426614174000
```

### Options

```
      --from-line int   Line from which to start retrieving logs
  -h, --help            help for build-logs
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

* [hostinger hosting nodejs](hostinger_hosting_nodejs.md)	 - NodeJS commands

