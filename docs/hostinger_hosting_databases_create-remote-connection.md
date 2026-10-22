## hostinger hosting databases create-remote-connection

Create database remote connection

### Synopsis

Allows a remote host to connect to the specified database.

Provide an IPv4/IPv6 address, or "%" to allow any host. The database name must be
the full name returned by the list databases endpoint. Database creation is synchronous,
so a 404 here means no database with that name exists under the username, not that it
is still being created.

```
hostinger hosting databases create-remote-connection <username> <name> [flags]
```

### Examples

```
  hostinger hosting databases create-remote-connection u123456789 u123456789_test_db --ip 192.0.2.10
```

### Options

```
  -h, --help        help for create-remote-connection
      --ip string   Remote host to allow: an IPv4/IPv6 address, or "%" for any host.
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

* [hostinger hosting databases](hostinger_hosting_databases.md)	 - Databases commands

