## hostinger hosting databases list-remote-connections

List database remote connections

### Synopsis

Returns the remote-access rules for the specified account: the remote hosts
(IPv4/IPv6 addresses, or "%" for any host) allowed to connect to the account databases.

Use the domain filter to only return rules for databases assigned to a specific domain.

```
hostinger hosting databases list-remote-connections <username> [flags]
```

### Examples

```
  hostinger hosting databases list-remote-connections u123456789
```

### Options

```
      --domain string   Filter remote connections by the domain the database is assigned to. Rules for databases not assigned to any domain are always included.
  -h, --help            help for list-remote-connections
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

