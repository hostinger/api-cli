## hostinger wordpress themes deploy

Deploy WordPress theme

### Synopsis

Deploy a WordPress theme from an already uploaded directory.

This endpoint allows you to deploy a WordPress theme that has been uploaded to the website's directory.
The theme can be optionally activated after deployment.

```
hostinger wordpress themes deploy <username> <domain> [flags]
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
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress themes](hostinger_wordpress_themes.md)	 - Themes commands

