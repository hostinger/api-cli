## hostinger hosting nodejs build-settings

Get Node.js build settings

### Synopsis

Returns the build settings stored for the website: framework (`app_type`), Node.js version,
root and output directory, build script, entry file and package manager. Stored settings
drive Git auto-deployment builds. A build started through the API uses the values sent in
that request and saves them here only when no settings exist yet.

Returns 404 until the first build or the first settings update stores them. Use this after
a failed build to check whether the framework or the entry file were detected wrong, then
fix them with the `Update Node.js build settings` endpoint.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting nodejs build-settings [domain] [flags]
```

### Examples

```
  hostinger hosting nodejs build-settings mydomain.tld
```

### Options

```
  -h, --help   help for build-settings
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

