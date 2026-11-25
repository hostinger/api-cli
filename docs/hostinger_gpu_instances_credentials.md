## hostinger gpu instances credentials

Get instance credentials

### Synopsis

Retrieve SSH connection credentials (hostname, port, username and password) of a specified GPU instance.

Credentials become available once the instance is provisioned; until then the endpoint responds with
`409 Conflict`, so retry later.

Use this endpoint to connect to a GPU instance over SSH.

```
hostinger gpu instances credentials <instance-id> [flags]
```

### Examples

```
  hostinger gpu instances credentials 6325
```

### Options

```
  -h, --help[=true|false]   help for credentials
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

* [hostinger gpu instances](hostinger_gpu_instances.md)	 - Instances commands

