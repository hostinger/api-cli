## hostinger hosting pocketbase create-backend

Create PocketBase backend

### Synopsis

Creates a PocketBase backend for a website: a database, authentication, file storage and a
realtime REST API, served from the website's own domain. A website has at most one
PocketBase backend. Creation usually takes a few seconds.

The backend takes over the `/api` and `/_` paths of the website: from then on requests to
these paths and everything below them go to PocketBase, while every other path keeps
serving the website. Do not create a backend for a website that serves its own `/api` or
`/_` paths, because they stop working; confirm with the user before calling it. Deleting
the backend gives both paths back to the website. Headless WordPress websites and parked
domains are not supported; for a parked domain, use the website it is parked on. A website
that already has a backend is rejected with 422, unless that backend is in state
`missing`: then it is replaced with a new, empty one. A website set up with
`type: headless_empty` becomes `headless_pocketbase`; any other website keeps its type.

`superuser_password` is optional: leave it out to get a strong generated password, or send
your own (10-128 characters). The response returns it in `superuser_password` once and
Hostinger never stores it, so save it in your app secrets right away; you need it to manage
the backend through the PocketBase REST API. If it is lost, set a new one with the
`Set PocketBase superuser` endpoint. `superuser_email` defaults to the email of the
website owner.

The instance sleeps when idle and wakes on the first request, so the first call after a
quiet period can take a few seconds. `next_steps` in the response lists what to do next.

PocketBase is rolling out gradually: an account without access yet gets 422.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting pocketbase create-backend [domain] [flags]
```

### Examples

```
  hostinger hosting pocketbase create-backend mydomain.tld
```

### Options

```
  -h, --help[=true|false]           help for create-backend
      --superuser-email string      Email of the PocketBase superuser. Defaults to the email of the website owner.
      --superuser-password string   Password for the PocketBase superuser, 10-128 characters. Leave it out to have a strong one generated. Either way it is returned once in the response and never stored.
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

