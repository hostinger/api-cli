## hostinger hosting php update-options

Update PHP options

### Synopsis

Updates PHP options for the website (e.g. `memory_limit`, `max_execution_time`, `upload_max_filesize`).
Only provide the options you want to change, inside the `options` object.

Values above the account plan limit are silently capped to that limit, so the request can succeed
with a smaller applied value. Call the Get PHP details endpoint afterwards to read the applied value.

The website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --options
    {}

```
hostinger hosting php update-options [domain] [flags]
```

### Examples

```
  hostinger hosting php update-options mydomain.tld --options '{}'
```

### Options

```
  -h, --help             help for update-options
      --options string   Map of PHP options to update, keyed by option name. Only include options you want to change. (JSON, @file or -; shape in the help above)
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

* [hostinger hosting php](hostinger_hosting_php.md)	 - PHP commands

