## hostinger gpu instances

Instances commands

```
hostinger gpu instances [flags]
```

### Options

```
  -h, --help[=true|false]   help for instances
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

* [hostinger gpu](hostinger_gpu.md)	 - GPU commands
* [hostinger gpu instances create](hostinger_gpu_instances_create.md)	 - Create instance
* [hostinger gpu instances delete](hostinger_gpu_instances_delete.md)	 - Delete instance
* [hostinger gpu instances get](hostinger_gpu_instances_get.md)	 - Get instance details
* [hostinger gpu instances list](hostinger_gpu_instances_list.md)	 - Get instances
* [hostinger gpu instances metrics](hostinger_gpu_instances_metrics.md)	 - Get instance metrics

