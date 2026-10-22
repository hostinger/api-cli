## hostinger wordpress themes deploy

Deploy WordPress theme

### Synopsis

Deploy a WordPress theme from an already uploaded directory.

This endpoint allows you to deploy a WordPress theme that has been uploaded to the website's directory.
The theme can be optionally activated after deployment.

```
hostinger wordpress themes deploy <username> <domain> [flags]
```

### Examples

```
  hostinger wordpress themes deploy u123456789 mydomain.tld --slug twentytwentyone --theme-path twentytwentyone-new
```

### Options

```
  -h, --help                help for deploy
      --is-activated        Whether to activate the theme after deployment
      --slug string         Slug of the theme
      --theme-path string   Relative path to the theme directory from wp-content/themes
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

* [hostinger wordpress themes](hostinger_wordpress_themes.md)	 - Themes commands

