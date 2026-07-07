## hostinger wordpress litespeed-cache purge-lite-speed

Purge LiteSpeed Cache

### Synopsis

Purge the LiteSpeed Cache for the specified WordPress installation.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

```
hostinger wordpress litespeed-cache purge-lite-speed <username> <software> [flags]
```

### Options

```
  -h, --help   help for purge-lite-speed
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress litespeed-cache](hostinger_wordpress_litespeed-cache.md)	 - LiteSpeed Cache commands

