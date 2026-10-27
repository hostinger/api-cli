## hostinger wordpress installations import-website

Import WordPress website

### Synopsis

Import WordPress website to the specified domain.

WARNING: this overwrites the website's existing contents and cannot be undone —
verify this is intended before calling this endpoint.

This endpoint allows you to import a WordPress website from archive and
database files that have been uploaded to the website's directory.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger wordpress installations import-website [domain] [flags]
```

### Examples

```
  hostinger wordpress installations import-website mydomain.tld --archive-path backup.zip --sql-path database.sql
```

### Options

```
      --archive-path string   Path to the WordPress archive file (relative to website root)
  -h, --help                  help for import-website
      --sql-path string       Path to the database SQL file (relative to website root)
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

* [hostinger wordpress installations](hostinger_wordpress_installations.md)	 - Installations commands

