## hostinger wordpress plugins deploy

Deploy WordPress plugin

### Synopsis

Deploy a WordPress plugin from an already uploaded directory.

This endpoint allows you to deploy a WordPress plugin that has been uploaded to the website's directory.
The plugin will be activated and made available in the WordPress admin panel.

```
hostinger wordpress plugins deploy <username> <domain> [flags]
```

### Examples

```
  hostinger wordpress plugins deploy u123456789 mydomain.tld --plugin-path my-plugin-new --slug my-plugin
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
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger wordpress plugins](hostinger_wordpress_plugins.md)	 - Plugins commands

