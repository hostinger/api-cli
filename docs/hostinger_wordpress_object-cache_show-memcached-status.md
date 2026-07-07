## hostinger wordpress object-cache show-memcached-status

Show Memcached object cache status

### Synopsis

Show the Memcached object cache status for the specified WordPress
installation.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

```
hostinger wordpress object-cache show-memcached-status <username> <software> [flags]
```

### Options

```
  -h, --help   help for show-memcached-status
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress object-cache](hostinger_wordpress_object-cache.md)	 - Object Cache commands

