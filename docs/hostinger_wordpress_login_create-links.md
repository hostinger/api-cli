## hostinger wordpress login create-links

Create login links

### Synopsis

Create temporary auto-login links for the specified WordPress installation.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

```
hostinger wordpress login create-links <username> <software> [flags]
```

### Options

```
  -h, --help   help for create-links
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress login](hostinger_wordpress_login.md)	 - Login commands

