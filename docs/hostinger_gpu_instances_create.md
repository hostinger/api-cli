## hostinger gpu instances create

Create instance

### Synopsis

Create a new GPU instance.

Pick `pool_instance_id` from the pool instance types list and `os_template_id` from the OS templates
available for that instance type. Your account must have enough credits to run the instance for its
first hour.

The instance is created asynchronously: a successful request returns immediately, and the instance
appears in the instance list while its `status` moves from `created` through `creating` to `running`.

Use this endpoint to provision GPU compute.

```
hostinger gpu instances create [flags]
```

### Examples

```
  hostinger gpu instances create --name my-instance-1 --os-template-id 23 --pool-instance-id 232
```

### Options

```
  -h, --help[=true|false]      help for create
      --name string            Instance name, unique among your instances. Letters, numbers, dashes and underscores only
      --os-template-id int     OS template ID available for the selected instance type
      --pool-instance-id int   Instance type ID from the pool instance types list
      --public-keys strings    SSH public keys to add to the instance
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

