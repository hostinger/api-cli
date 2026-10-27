## hostinger horizons websites list

Get website list

### Synopsis

List the Hostinger Horizons websites the user owns.\n
Use this tool when the user asks which websites they have, or when you need a website ID
before editing, publishing or cloning a website.\n
Each website is returned with its ID, status, domain and the URL to open it
in Hostinger Horizons interface.\n
The complete list of websites is returned in a single response - it is not paginated.

```
hostinger horizons websites list [flags]
```

### Examples

```
  hostinger horizons websites list
```

### Options

```
  -h, --help   help for list
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

* [hostinger horizons websites](hostinger_horizons_websites.md)	 - Websites commands

