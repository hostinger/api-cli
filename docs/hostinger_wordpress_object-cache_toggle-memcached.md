## hostinger wordpress object-cache toggle-memcached

Toggle Memcached object cache

### Synopsis

Activate or deactivate the Memcached object cache for the specified WordPress
installation, based on the `enabled` flag.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

```
hostinger wordpress object-cache toggle-memcached <username> <software> [flags]
```

### Options

```
      --enabled   Activate (true) or deactivate (false) the Memcached object cache for the WordPress installation.
  -h, --help      help for toggle-memcached
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress object-cache](hostinger_wordpress_object-cache.md)	 - Object Cache commands

