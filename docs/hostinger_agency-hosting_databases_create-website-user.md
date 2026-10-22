## hostinger agency-hosting databases create-website-user

Create website database user

### Synopsis

Creates a user for an existing database on an Agency Plan website.

Each database supports a single non-system user; creating a user for a database that already has one fails.

```
hostinger agency-hosting databases create-website-user <website_uid> <database_name> [flags]
```

### Examples

```
  hostinger agency-hosting databases create-website-user zpwlGlp19 my_database --database-user my_user --password 'Str0ngP@ssword!'
```

### Options

```
      --database-user string   Database username to create (alphanumeric and underscores).
  -h, --help                   help for create-website-user
      --host string            Host the user connects from (IPv4, IPv6, % wildcard, or localhost). Defaults to localhost.
      --password string        Password for the database user (requires mixed case, letters, and numbers).
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

