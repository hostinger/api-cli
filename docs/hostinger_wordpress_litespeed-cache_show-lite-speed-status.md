## hostinger wordpress litespeed-cache show-lite-speed-status

Show LiteSpeed Cache status

### Synopsis

Show the LiteSpeed Cache status for the specified WordPress installation.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

```
hostinger wordpress litespeed-cache show-lite-speed-status <username> <software> [flags]
```

### Options

```
  -h, --help   help for show-lite-speed-status
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress litespeed-cache](hostinger_wordpress_litespeed-cache.md)	 - LiteSpeed Cache commands

