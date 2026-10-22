## hostinger wordpress plugins check-if-woo-commerce-is-installed

Check if WooCommerce is installed

### Synopsis

Check whether WooCommerce is installed on any WordPress installation of a
domain. Optionally filter by domain to scope the check.

```
hostinger wordpress plugins check-if-woo-commerce-is-installed [flags]
```

### Examples

```
  hostinger wordpress plugins check-if-woo-commerce-is-installed
```

### Options

```
      --domain string   Filter by domain name (case-insensitive substring match)
  -h, --help            help for check-if-woo-commerce-is-installed
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

