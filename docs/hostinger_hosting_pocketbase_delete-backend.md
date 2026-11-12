## hostinger hosting pocketbase delete-backend

Delete PocketBase backend

### Synopsis

Permanently deletes the PocketBase backend of a website with all its data (collections,
records, users and uploaded files) and gives the `/api` and `/_` paths back to the website;
a `headless_pocketbase` website becomes `headless_empty`. This cannot be undone and every
app using this backend stops working, so confirm with the user before calling it and
export the data you need from the dashboard (`admin_url`) first.

Returns success also when the website has no PocketBase backend; then it only clears a
`headless_pocketbase` mark. To regain access after losing the superuser password, use the
`Set PocketBase superuser` endpoint instead of deleting the backend.

PocketBase is rolling out gradually: an account without access yet gets 422; an unknown
website gets 404.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting pocketbase delete-backend [domain] [flags]
```

### Examples

```
  hostinger hosting pocketbase delete-backend mydomain.tld
```

### Options

```
  -h, --help[=true|false]   help for delete-backend
  -y, --yes[=true|false]    Skip the confirmation prompt
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

* [hostinger hosting pocketbase](hostinger_hosting_pocketbase.md)	 - PocketBase commands

