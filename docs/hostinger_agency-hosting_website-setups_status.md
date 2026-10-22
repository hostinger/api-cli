## hostinger agency-hosting website-setups status

Get website setup status

### Synopsis

Returns the current status of an Agency Plan website setup started via the setups
endpoint.

Poll this endpoint using the `setup_uuid` returned from the provisioning request until
`status` becomes `completed`, at which point `website_uid` identifies the new website.

```
hostinger agency-hosting website-setups status <order_id> <setup_uuid> [flags]
```

### Examples

```
  hostinger agency-hosting website-setups status 123456 0193b6d4-fabb-70e0-8ea4-cfe060a45898
```

### Options

```
  -h, --help   help for status
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

* [hostinger agency-hosting website-setups](hostinger_agency-hosting_website-setups.md)	 - Website Setups commands

