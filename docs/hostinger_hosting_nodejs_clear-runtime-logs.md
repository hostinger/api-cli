## hostinger hosting nodejs clear-runtime-logs

Clear Node.js runtime logs

### Synopsis

Empties the Node.js application's runtime log file. This cannot be undone, so confirm with
the user before calling it. Returns success even when no log file exists yet.

Use it before reproducing a problem so the next `Get Node.js runtime logs` call returns
only fresh entries; start that call with `period` again instead of reusing a `from_line`
from before the clear.

```
hostinger hosting nodejs clear-runtime-logs <username> <domain> [flags]
```

### Examples

```
  hostinger hosting nodejs clear-runtime-logs u123456789 mydomain.tld
```

### Options

```
  -h, --help   help for clear-runtime-logs
  -y, --yes    Skip the confirmation prompt
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

