## hostinger hosting databases create-remote-connection

Create database remote connection

### Synopsis

Allows a remote host to connect to the specified database.

Provide an IPv4/IPv6 address, or "%" to allow any host. The database name must be
the full name returned by the list databases endpoint. Database creation is synchronous,
so a 404 here means no database with that name exists under the username, not that it
is still being created.

The hosting account can be given as [username] or left out, in which case it comes from --username, HOSTINGER_USERNAME, the website named by --domain or HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting databases create-remote-connection [username] <name> [flags]
```

### Examples

```
  hostinger hosting databases create-remote-connection u123456789 u123456789_test_db --ip 192.0.2.10
```

### Options

```
  -h, --help        help for create-remote-connection
      --ip string   Remote host to allow: an IPv4/IPv6 address, or "%" for any host.
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

