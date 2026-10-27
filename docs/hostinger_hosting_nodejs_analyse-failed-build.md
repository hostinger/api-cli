## hostinger hosting nodejs analyse-failed-build

Analyse failed Node.js build

### Synopsis

Returns an AI analysis of why a build failed and how to fix it, based on the build logs,
the project file list and package.json. Only builds in the `failed` state can be analysed;
any other state returns 422. When no analysis could be produced both `analysis` and
`solution` are null, in which case read `Get NodeJS build logs` instead.

Each call runs the analysis again, so call it once per failed build and keep the result.
Limited to 5 calls per minute per API client (429 above that).

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

The build can be given as [build-uuid] or left out, in which case it is the build `hostinger up` last started for the website (remembered in .hostinger/site.json), or a choice among the website's builds when run in a terminal.

```
hostinger hosting nodejs analyse-failed-build [domain] [build-uuid] [flags]
```

### Examples

```
  hostinger hosting nodejs analyse-failed-build mydomain.tld 123e4567-e89b-12d3-a456-426614174000
```

### Options

```
  -h, --help   help for analyse-failed-build
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

* [hostinger hosting nodejs](hostinger_hosting_nodejs.md)	 - NodeJS commands

