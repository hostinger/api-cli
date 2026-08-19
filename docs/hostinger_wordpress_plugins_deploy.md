## hostinger wordpress plugins deploy

Deploy WordPress plugin

### Synopsis

Deploy a WordPress plugin from an already uploaded directory.

This endpoint allows you to deploy a WordPress plugin that has been uploaded to the website's directory.
The plugin will be activated and made available in the WordPress admin panel.

```
hostinger wordpress plugins deploy <username> <domain> [flags]
```

### Options

```
  -h, --help                 help for deploy
      --plugin-path string   Relative path to the plugin directory from wp-content/plugins
      --slug string          Slug of the plugin
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress plugins](hostinger_wordpress_plugins.md)	 - Plugins commands

