## hostinger hosting databases delete-remote-connection

Delete account database remote connection

### Synopsis

Permanently removes a remote-access rule, revoking the given host's remote access to the database.

Identify the rule with the required ip query parameter (the IPv4/IPv6 address, or "%",
exactly as returned by the list remote connections endpoint). The database name must be
the full name returned by the list databases endpoint.

```
hostinger hosting databases delete-remote-connection <username> <name> [flags]
```

### Options

```
  -h, --help        help for delete-remote-connection
      --ip string   Remote host to revoke: the IPv4/IPv6 address, or "%",
                    exactly as returned by the list remote connections endpoint.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting databases](hostinger_hosting_databases.md)	 - Databases commands

