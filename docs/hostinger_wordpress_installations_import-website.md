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

### Options

```
      --archive-path string   Path to the WordPress archive file (relative to website root)
  -h, --help                  help for import-website
      --sql-path string       Path to the database SQL file (relative to website root)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress installations](hostinger_wordpress_installations.md)	 - Installations commands

