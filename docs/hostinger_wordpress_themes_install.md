## hostinger wordpress themes install

Install WordPress theme

### Synopsis

Install a theme on an existing WordPress installation.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id`
field).

When the theme is one of the Hostinger themes (hostinger-blog,
hostinger-affiliate-theme, hostinger-ai-theme), the optional `palette`,
`layout`, and `font` fields are forwarded to the custom installer (defaults:
palette1, layout1, default). For any other theme they are ignored.

This operation is asynchronous: a successful response only means the install
job has been queued, not that the theme is ready.

```
hostinger wordpress themes install <username> <software> [flags]
```

### Options

```
      --font string      Font identifier. Only applied when the theme is a Hostinger theme; the default is used when omitted. (one of: professional, modern, elegant, creative, dynamic, default) (default "default")
  -h, --help             help for install
      --layout string    Layout identifier. Only applied when the theme is a Hostinger theme; the default is used when omitted. (default "layout1")
      --palette string   Palette identifier. Only applied when the theme is a Hostinger theme; the default is used when omitted. (default "palette1")
      --theme string     Slug of the theme to install. Hostinger theme slugs (hostinger-blog, hostinger-affiliate-theme, hostinger-ai-theme) trigger the custom installer and forward the optional palette/layout/font fields; any other WordPress theme slug uses the standard installer and ignores those fields.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress themes](hostinger_wordpress_themes.md)	 - Themes commands

