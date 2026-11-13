## hostinger hosting pocketbase backend

Get PocketBase backend

### Synopsis

Returns the PocketBase backend of a website: its REST API and dashboard URLs, live state
and version. A PocketBase backend gives the website a database, authentication, file
storage and a realtime REST API, served from the website's own domain under `/api` and
`/_`. Returns 404 when the website has no PocketBase backend.

The superuser password is only shown once, when it is set, and is never returned here. If it
is lost, set a new one with the `Set PocketBase superuser` endpoint.

The instance sleeps when idle and wakes on the first request, so `asleep` is a normal state
and the first call to a sleeping instance can take a few seconds. `starting`, `running` and
`stopping` are healthy.

`failed` means the last start did not succeed. Hostinger starts it again on its own after
`retry_after_seconds`, and the data is intact: wait and call again, do not delete the
backend. `failure_reason` says why: `process_limit` means the hosting account has no free
process slots, usually because its Node.js apps are running.

`unregistered` means the instance is broken. `missing` means its files are gone from the
server, and creating the PocketBase backend again sets up a new, empty one. `state` is null
when the server could not be reached; retry later.

PocketBase is rolling out gradually: an account without access yet gets 422.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting pocketbase backend [domain] [flags]
```

### Examples

```
  hostinger hosting pocketbase backend mydomain.tld
```

### Options

```
  -h, --help[=true|false]   help for backend
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

