## hostinger agency-hosting php replace-website-options

Replace website PHP options

### Synopsis

Replaces the custom php.ini values on an Agency Plan website with the ones provided. Any option not in the request is reset to its default, so call the options endpoint first and send the full desired set. Sending an empty array resets every option to its default.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --options
    [
      {
        "name": "memory_limit",
        "value": "256M"
      }
    ]

```
hostinger agency-hosting php replace-website-options <website_uid> [flags]
```

### Examples

```
  hostinger agency-hosting php replace-website-options zpwlGlp19 --options '[{"name":"memory_limit","value":"256M"}]'
```

### Options

```
  -h, --help             help for replace-website-options
      --options string   Option names and values. Each name must be one of the options returned by the options endpoint, and each value must satisfy that option's allowed_values when it declares them. (JSON, @file or -; shape in the help above)
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

* [hostinger agency-hosting php](hostinger_agency-hosting_php.md)	 - PHP commands

