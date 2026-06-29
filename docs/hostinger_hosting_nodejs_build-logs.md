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

### Options

```
      --from-line int   Line from which to start retrieving logs
  -h, --help            help for build-logs
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting nodejs](hostinger_hosting_nodejs.md)	 - NodeJS commands

