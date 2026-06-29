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

### Options

```
  -h, --help             help for list-builds
      --page int         Page number
      --per-page int     Number of items per page (default 25)
      --states strings   Build states to filter by (one of: pending, running, completed, failed)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting nodejs](hostinger_hosting_nodejs.md)	 - NodeJS commands

