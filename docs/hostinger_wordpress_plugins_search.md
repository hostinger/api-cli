## hostinger wordpress plugins search

Search WordPress plugins

### Synopsis

Search the WordPress.org plugin directory for plugins available to install.

Use the returned `slug` values with
POST /api/hosting/v1/accounts/{username}/wordpress/{software}/plugins/install.

```
hostinger wordpress plugins search [flags]
```

### Options

```
  -h, --help            help for search
      --search string   Search term to match against plugin names. Minimum 3 characters.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress plugins](hostinger_wordpress_plugins.md)	 - Plugins commands

