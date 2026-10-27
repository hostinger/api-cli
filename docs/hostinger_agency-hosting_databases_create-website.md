## hostinger agency-hosting databases create-website

Create website database

### Synopsis

Creates a MySQL database with a dedicated user for an Agency Plan website.

The database name, username, and password must all be provided by the caller.

```
hostinger agency-hosting databases create-website <website_uid> [flags]
```

### Examples

```
  hostinger agency-hosting databases create-website zpwlGlp19 --database-name mydatabase --database-user myuser --password 'Str0ngP@ssword!'
```

### Options

```
      --database-name string   Database name to create (alphanumeric characters).
      --database-user string   Database username to create alongside the database (alphanumeric characters).
  -h, --help                   help for create-website
      --password string        Password for the database user (requires mixed case, letters, and numbers).
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

* [hostinger agency-hosting databases](hostinger_agency-hosting_databases.md)	 - Databases commands

