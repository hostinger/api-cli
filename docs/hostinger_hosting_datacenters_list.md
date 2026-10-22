## hostinger hosting datacenters list

List available datacenters

### Synopsis

Retrieve a list of datacenters available for setting up hosting plans
based on available datacenter capacity and hosting plan of your order.
The first item in the list is the best match for your specific order
requirements.

```
hostinger hosting datacenters list [flags]
```

### Examples

```
  hostinger hosting datacenters list --order-id 123
```

### Options

```
  -h, --help           help for list
      --order-id int   Order ID
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger hosting datacenters](hostinger_hosting_datacenters.md)	 - Datacenters commands

