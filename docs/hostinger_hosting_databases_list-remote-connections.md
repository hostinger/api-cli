## hostinger hosting databases list-remote-connections

List account database remote connections

### Synopsis

Returns the remote-access rules for the specified account: the remote hosts
(IPv4/IPv6 addresses, or "%" for any host) allowed to connect to the account databases.

Use the domain filter to only return rules for databases assigned to a specific domain.

```
hostinger hosting databases list-remote-connections <username> [flags]
```

### Options

```
      --domain string   Filter remote connections by the domain the database is assigned to.
                        Rules for databases not assigned to any domain are always included.
  -h, --help            help for list-remote-connections
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting databases](hostinger_hosting_databases.md)	 - Databases commands

