## hostinger hosting websites deploy-static-site-archive

Deploy static site archive

### Synopsis

Deploy a static application from an archive file.

WARNING: this overwrites the website's existing contents and cannot be undone —
verify this is intended before calling this endpoint.

This endpoint allows you to deploy a static application from an archive
file that has been uploaded to the website's directory.

This only works for static sites (pre-built HTML/CSS/JS with no build step). For
Node.js applications, use `Create NodeJS build from archive` instead, or
`Start Node.js build` if the archive is already uploaded. For WordPress sites,
use `Import WordPress website`.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting websites deploy-static-site-archive [domain] [flags]
```

### Examples

```
  hostinger hosting websites deploy-static-site-archive mydomain.tld --archive-path uploads/archive.zip
```

### Options

```
      --archive-path string   Relative path to the archive file from website root directory
  -h, --help                  help for deploy-static-site-archive
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

* [hostinger hosting websites](hostinger_hosting_websites.md)	 - Websites commands

