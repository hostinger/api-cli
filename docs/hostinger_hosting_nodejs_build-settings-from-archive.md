## hostinger hosting nodejs build-settings-from-archive

Get Node.js build settings from archive

### Synopsis

Auto-detect Node.js build settings from a package.json inside an archive already on the server.

Use this before calling `Start Node.js Build` to preview what settings will be used,
or to let the user review and override values (framework, node version, root directory,
output directory, build script) before committing to a build.

The archive must already be present on the website's file storage. Use the
`Generate Upload URL` endpoint to obtain credentials and upload the archive first.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

```
hostinger hosting nodejs build-settings-from-archive [domain] [flags]
```

### Examples

```
  hostinger hosting nodejs build-settings-from-archive mydomain.tld --archive-path value
```

### Options

```
      --archive-path string   The path to the archive file relative to the document root of the vhost
  -h, --help                  help for build-settings-from-archive
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

