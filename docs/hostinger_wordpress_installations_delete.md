## hostinger wordpress installations delete

Delete WordPress installation

### Synopsis

Delete the specified WordPress installation, with optional file and database
removal. This removes all associated components including plugins, themes,
staging websites and any other related data.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

```
hostinger wordpress installations delete <username> <software> [flags]
```

### Examples

```
  hostinger wordpress installations delete u123456789 1232456789
```

### Options

```
      --delete-database   Delete the installation database.
      --delete-files      Delete installation files from disk.
  -h, --help              help for delete
  -y, --yes               Skip the confirmation prompt
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

