## hostinger hosting databases create

Create account database

### Synopsis

Creates a database with a database user and password for the specified account.

The database name and user are automatically prefixed with the account username when needed.

The hosting account can be given as [username] or left out, in which case it comes from --username, HOSTINGER_USERNAME, the website named by --domain or HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting databases create [username] [flags]
```

### Examples

```
  hostinger hosting databases create u123456789 --name u123456789_test_db --password 'Str0ngP@ssword!' --user u123456789_admin --website-domain example.com
```

### Options

```
  -h, --help                    help for create
      --name string             Database name. If the account username prefix is omitted, it is added automatically.
      --password string         Database user password.
      --user string             Database user. If the account username prefix is omitted, it is added automatically.
      --website-domain string   Website domain assigned to the database.
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

* [hostinger hosting databases](hostinger_hosting_databases.md)	 - Databases commands

