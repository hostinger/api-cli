## hostinger hosting databases delete-remote-connection

Delete database remote connection

### Synopsis

Permanently removes a remote-access rule, revoking the given host's remote access to the database.

Identify the rule with the required ip query parameter (the IPv4/IPv6 address, or "%",
exactly as returned by the list remote connections endpoint). The database name must be
the full name returned by the list databases endpoint.

```
hostinger hosting databases delete-remote-connection <username> <name> [flags]
```

### Examples

```
  hostinger hosting databases delete-remote-connection u123456789 u123456789_test_db --ip 192.0.2.10
```

### Options

```
  -h, --help        help for delete-remote-connection
      --ip string   Remote host to revoke: the IPv4/IPv6 address, or "%", exactly as returned by the list remote connections endpoint.
  -y, --yes         Skip the confirmation prompt
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

