## hostinger hosting databases create-remote-connection

Create account database remote connection

### Synopsis

Allows a remote host to connect to the specified database.

Provide an IPv4/IPv6 address, or "%" to allow any host. The database name must be
the full name returned by the list databases endpoint.

```
hostinger hosting databases create-remote-connection <username> <name> [flags]
```

### Options

```
  -h, --help        help for create-remote-connection
      --ip string   Remote host to allow: an IPv4/IPv6 address, or "%" for any host.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting databases](hostinger_hosting_databases.md)	 - Databases commands

