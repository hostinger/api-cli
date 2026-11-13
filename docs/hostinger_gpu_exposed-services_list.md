## hostinger gpu exposed-services list

Get exposed services

### Synopsis

Retrieve the services exposed on a specified GPU instance.

A newly created service is listed once it gets its public hostname and external port.

Use this endpoint to find the public address of services running on your GPU instance.

```
hostinger gpu exposed-services list <instance-id> [flags]
```

### Examples

```
  hostinger gpu exposed-services list 6325
```

### Options

```
  -h, --help[=true|false]   help for list
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

* [hostinger gpu exposed-services](hostinger_gpu_exposed-services.md)	 - Exposed services commands

