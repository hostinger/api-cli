## hostinger wordpress plugins search

Search WordPress plugins

### Synopsis

Search the WordPress.org plugin directory for plugins available to install.

Use the returned `slug` values with
POST /api/hosting/v1/accounts/{username}/wordpress/{software}/plugins/install.

```
hostinger wordpress plugins search [flags]
```

### Examples

```
  hostinger wordpress plugins search --search seo
```

### Options

```
  -h, --help            help for search
      --search string   Search term to match against plugin names. Minimum 3 characters.
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

