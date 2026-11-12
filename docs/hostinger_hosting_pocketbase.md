## hostinger hosting pocketbase

PocketBase commands

```
hostinger hosting pocketbase [flags]
```

### Options

```
  -h, --help[=true|false]   help for pocketbase
```

### Options inherited from parent commands

```
      --config string           Config file (default is $HOME/.hostinger.yaml)
      --domain string           Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run[=true|false]    Print the request as a curl command instead of sending it
  -o, --format string           Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string               Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input[=true|false]   Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string         Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string         Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose[=true|false]    Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger hosting](hostinger_hosting.md)	 - Hosting commands
* [hostinger hosting pocketbase backend](hostinger_hosting_pocketbase_backend.md)	 - Get PocketBase backend
* [hostinger hosting pocketbase create-backend](hostinger_hosting_pocketbase_create-backend.md)	 - Create PocketBase backend
* [hostinger hosting pocketbase delete-backend](hostinger_hosting_pocketbase_delete-backend.md)	 - Delete PocketBase backend
* [hostinger hosting pocketbase list-backends](hostinger_hosting_pocketbase_list-backends.md)	 - List account PocketBase backends
* [hostinger hosting pocketbase set-superuser](hostinger_hosting_pocketbase_set-superuser.md)	 - Set PocketBase superuser

