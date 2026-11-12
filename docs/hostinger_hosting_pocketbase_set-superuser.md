## hostinger hosting pocketbase set-superuser

Set PocketBase superuser

### Synopsis

Sets the password of a superuser of the website's PocketBase backend, and creates that
superuser when there is none with the given email. The old password of that superuser stops
working at once, so apps and scripts still using it lose superuser access; confirm with the
user before calling it.

Use it when the superuser password is lost (Hostinger never stores it), when it leaked and
must be rotated, or to add another superuser with a different email. Returns 404 when the
website has no PocketBase backend; create one first. App user accounts are managed through
the PocketBase REST API, not here.

`superuser_password` is optional: leave it out to get a strong generated password, or send
your own (10-128 characters). The response returns it in `superuser_password` once and
Hostinger never stores it, so save it in your app secrets right away, then authenticate
again with `pb.collection('_superusers').authWithPassword`. `superuser_email` defaults to
the email of the website owner.

PocketBase is rolling out gradually: an account without access yet gets 422.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting pocketbase set-superuser [domain] [flags]
```

### Examples

```
  hostinger hosting pocketbase set-superuser mydomain.tld
```

### Options

```
  -h, --help[=true|false]           help for set-superuser
      --superuser-email string      Email of the superuser to create or update. Defaults to the email of the website owner.
      --superuser-password string   New password of the PocketBase superuser, 10-128 characters. Leave it out to have a strong one generated. Either way it is returned once in the response and never stored.
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

