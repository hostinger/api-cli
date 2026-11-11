## hostinger gpu catalog pool-instance-types

Get pool instance types

### Synopsis

Retrieve the instance types offered by a specified GPU pool, ordered by CPU core count.

Each instance type describes its CPU, memory and storage, and its hourly price.
Only instance types that can currently be ordered are returned.

Use this endpoint to choose the `pool_instance_id` for a new GPU instance.

```
hostinger gpu catalog pool-instance-types <pool-id> [flags]
```

### Examples

```
  hostinger gpu catalog pool-instance-types 12
```

### Options

```
  -h, --help[=true|false]   help for pool-instance-types
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

* [hostinger gpu catalog](hostinger_gpu_catalog.md)	 - Catalog commands

