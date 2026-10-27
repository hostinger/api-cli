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

The hosting account can be given as [username] or left out, in which case it comes from --username, HOSTINGER_USERNAME, the website named by --domain or HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, default_domain in the config file, or a prompt when run in a terminal.

```
hostinger wordpress themes install [username] <software> [flags]
```

### Examples

```
  hostinger wordpress themes install u123456789 1232456789 --theme hostinger-blog
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
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --domain string     Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger wordpress themes](hostinger_wordpress_themes.md)	 - Themes commands

