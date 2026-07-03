## hostinger wordpress plugins list-suggested

List suggested WordPress plugins

### Synopsis

List curated plugin suggestions grouped by website type.

Use the returned `slug` values with
POST /api/hosting/v1/accounts/{username}/wordpress/{software}/plugins/install.

```
hostinger wordpress plugins list-suggested [flags]
```

### Options

```
  -h, --help           help for list-suggested
      --order-id int   Optionally scope suggestions to a specific order.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress plugins](hostinger_wordpress_plugins.md)	 - Plugins commands

