## hostinger wordpress installations import-website

Import WordPress website

### Synopsis

Import WordPress website to the specified domain.

WARNING: this overwrites the website's existing contents and cannot be undone —
verify this is intended before calling this endpoint.

This endpoint allows you to import a WordPress website from archive and
database files that have been uploaded to the website's directory.

```
hostinger wordpress installations import-website <username> <domain> [flags]
```

### Examples

```
  hostinger wordpress installations import-website u123456789 mydomain.tld --archive-path backup.zip --sql-path database.sql
```

### Options

```
      --archive-path string   Path to the WordPress archive file (relative to website root)
  -h, --help                  help for import-website
      --sql-path string       Path to the database SQL file (relative to website root)
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

* [hostinger wordpress installations](hostinger_wordpress_installations.md)	 - Installations commands

