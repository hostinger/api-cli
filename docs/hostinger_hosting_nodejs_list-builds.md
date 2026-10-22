## hostinger hosting nodejs list-builds

List NodeJS builds

### Synopsis

Retrieve a paginated list of Node.js build processes for a specific website.

Each build represents a single run of the Node.js build pipeline. Use the `states`
query parameter to filter results by build state (pending, running, completed, failed).
Use the `uuid` from a build to poll its output via the `Get Node.js Build Logs` endpoint.

```
hostinger hosting nodejs list-builds <username> <domain> [flags]
```

### Examples

```
  hostinger hosting nodejs list-builds u123456789 mydomain.tld
```

### Options

```
      --all              Fetch every page and print all items
  -h, --help             help for list-builds
      --page int         Page number
      --per-page int     Number of items per page (default 25)
      --states strings   Build states to filter by (one of: pending, running, completed, failed)
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

