## hostinger agency-hosting databases delete-website

Delete website database

### Synopsis

Permanently deletes a MySQL database and all its data from an Agency Plan website, including its users.

The operation is idempotent: deleting a database that does not exist succeeds without error.

```
hostinger agency-hosting databases delete-website <website_uid> <database_name> [flags]
```

### Examples

```
  hostinger agency-hosting databases delete-website zpwlGlp19 my_database
```

### Options

```
  -h, --help   help for delete-website
  -y, --yes    Skip the confirmation prompt
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

* [hostinger agency-hosting databases](hostinger_agency-hosting_databases.md)	 - Databases commands

