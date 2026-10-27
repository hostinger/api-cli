## hostinger hosting databases

Databases commands

```
hostinger hosting databases [flags]
```

### Options

```
  -h, --help   help for databases
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

* [hostinger hosting](hostinger_hosting.md)	 - Hosting commands
* [hostinger hosting databases change-password](hostinger_hosting_databases_change-password.md)	 - Change database password
* [hostinger hosting databases create](hostinger_hosting_databases_create.md)	 - Create account database
* [hostinger hosting databases create-remote-connection](hostinger_hosting_databases_create-remote-connection.md)	 - Create database remote connection
* [hostinger hosting databases delete](hostinger_hosting_databases_delete.md)	 - Delete account database
* [hostinger hosting databases delete-remote-connection](hostinger_hosting_databases_delete-remote-connection.md)	 - Delete database remote connection
* [hostinger hosting databases list](hostinger_hosting_databases_list.md)	 - List account databases
* [hostinger hosting databases list-remote-connections](hostinger_hosting_databases_list-remote-connections.md)	 - List database remote connections
* [hostinger hosting databases phpmyadmin-link](hostinger_hosting_databases_phpmyadmin-link.md)	 - Get phpMyAdmin link
* [hostinger hosting databases repair](hostinger_hosting_databases_repair.md)	 - Repair database
* [hostinger hosting databases setup-website](hostinger_hosting_databases_setup-website.md)	 - Setup website database

