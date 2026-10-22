## hostinger agency-hosting databases delete-website-user

Delete website database user

### Synopsis

Permanently deletes a database user from an Agency Plan website database, revoking all access it had.

The operation is idempotent: deleting a user that does not exist succeeds without error.

```
hostinger agency-hosting databases delete-website-user <website_uid> <database_name> <database_user_name> [flags]
```

### Examples

```
  hostinger agency-hosting databases delete-website-user zpwlGlp19 my_database my_user
```

### Options

```
  -h, --help   help for delete-website-user
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

